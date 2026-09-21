#!/usr/bin/env python3
"""Wire Identity hooks to repositories, rename files, fix unified-query imports."""
from __future__ import annotations

import re
from pathlib import Path

ROOT = Path("/volume1/Projects/NebulaForgeX/NFX-Identity/console")
SRC = ROOT / "src"

HOOK_MAP = {
    "useAuth.ts": ("useAuthRepository", "auth", "auth.ts"),
    "useDirectory.ts": ("useDirectoryRepository", "directory", "directory.ts"),
    "useAccess.ts": ("useAccessRepository", "access", "access.ts"),
    "useAudit.ts": ("useAuditRepository", "audit", "audit.ts"),
    "useClients.ts": ("useClientsRepository", "clients", "clients.ts"),
    "useImage.ts": ("useImageRepository", "image", "image.ts"),
    "useSystem.ts": ("useSystemRepository", "system", "system.ts"),
    "useTenants.ts": ("useTenantsRepository", "tenants", "tenants.ts"),
}


def extract_value_imports(block: str) -> list[str]:
    names: list[str] = []
    for raw in block.split(","):
        raw = raw.strip()
        if not raw or raw.startswith("type "):
            continue
        names.append(raw.split(" as ")[0].strip())
    return names


def convert_hook(path: Path, hook_fn: str, alias: str) -> None:
    text = path.read_text()
    m = re.search(
        r'import \{\n?(.*?)\n\} from "@/apis/\w+\.api";',
        text,
        re.S,
    )
    if not m:
        print("no api import", path)
        return
    imported = extract_value_imports(m.group(1))
    text = text[: m.start()] + f'import {{ {hook_fn} }} from "@/apis/repositories";' + text[m.end() :]

    def insert_repo(match: re.Match[str]) -> str:
        body = match.group(0)
        if f"const {alias} = {hook_fn}()" in body:
            return body
        return match.group(1) + f"\n  const {alias} = {hook_fn}();"

    text = re.sub(
        r"(export const use\w+ = \([^)]*\) => \{)",
        insert_repo,
        text,
    )

    for name in sorted(imported, key=len, reverse=True):
        text = re.sub(rf"\b{name}\(", f"{alias}.{name}(", text)

    used = set(re.findall(r"useUnified(?:Suspense)?(?:Infinite)?Query", text))
    type_bits = []
    if "UnifiedQueryParams" in text:
        type_bits.append("type UnifiedQueryParams")
    if "NormalUnifiedQueryOptions" in text:
        type_bits.append("type NormalUnifiedQueryOptions")
    hooks = []
    for h in (
        "useUnifiedQuery",
        "useUnifiedSuspenseQuery",
        "useUnifiedInfiniteQuery",
        "useUnifiedSuspenseInfiniteQuery",
    ):
        if h in used:
            hooks.append(h)
    import_items = hooks + type_bits
    if import_items:
        text = re.sub(
            r'import \{[^}]*\} from "nfx-ui/hooks";',
            "import { " + ", ".join(import_items) + ' } from "nfx-ui/hooks";',
            text,
            count=1,
        )
        if 'from "nfx-ui/hooks"' not in text.split("import { " + ", ".join(import_items))[0][-80:] if False else True:
            if not re.search(r'from "nfx-ui/hooks"', text):
                text = text.replace(
                    f'import {{ {hook_fn} }} from "@/apis/repositories";',
                    f'import {{ {hook_fn} }} from "@/apis/repositories";\nimport {{ {", ".join(import_items)} }} from "nfx-ui/hooks";',
                )

    path.write_text(text)
    print("wired", path.name)


def rename_hooks() -> None:
    hooks_dir = SRC / "hooks"
    for old, (_fn, _alias, new) in HOOK_MAP.items():
        src = hooks_dir / old
        dst = hooks_dir / new
        if src.exists():
            src.rename(dst)
            print("renamed", old, "->", new)

    index = hooks_dir / "index.ts"
    index.write_text(
        """export * from "./auth";
export * from "./resendTimer";
export * from "./access";
export * from "./audit";
export * from "./clients";
export * from "./directory";
export * from "./image";
export * from "./tenants";
"""
    )

    resend = hooks_dir / "useResendTimer.ts"
    if resend.exists():
        resend.rename(hooks_dir / "resendTimer.ts")

    # rewrite imports across src
    replacements = {
        "@/hooks/useAuth": "@/hooks/auth",
        "@/hooks/useDirectory": "@/hooks/directory",
        "@/hooks/useAccess": "@/hooks/access",
        "@/hooks/useAudit": "@/hooks/audit",
        "@/hooks/useClients": "@/hooks/clients",
        "@/hooks/useImage": "@/hooks/image",
        "@/hooks/useTenants": "@/hooks/tenants",
        "@/hooks/useResendTimer": "@/hooks/resendTimer",
        "@/hooks/useStyles": "@/hooks/styles",
        "@/hooks/useUserPreferenceSync": "@/hooks/preferenceSync",
    }
    styles = hooks_dir / "useStyles.ts"
    if styles.exists():
        styles.rename(hooks_dir / "styles.ts")

    for path in list(SRC.rglob("*.ts")) + list(SRC.rglob("*.tsx")):
        text = path.read_text()
        orig = text
        for a, b in replacements.items():
            text = text.replace(a, b)
        if text != orig:
            path.write_text(text)
            print("import rewrite", path.relative_to(ROOT))


def rename_inv() -> None:
    qdir = SRC / "providers/QueryProvider/hooks"
    mapping = {
        "useAuthCacheInvalidation.ts": "useAuthInv.ts",
        "useAccessCacheInvalidation.ts": "useAccessInv.ts",
        "useAuditCacheInvalidation.ts": "useAuditInv.ts",
        "useClientsCacheInvalidation.ts": "useClientsInv.ts",
        "useDirectoryCacheInvalidation.ts": "useDirectoryInv.ts",
        "useImageCacheInvalidation.ts": "useImageInv.ts",
        "useSystemCacheInvalidation.ts": "useSystemInv.ts",
        "useTenantsCacheInvalidation.ts": "useTenantsInv.ts",
        "useCacheInvalidationEvents.ts": "useQueryInv.ts",
    }
    fn_map = {
        "useAuthCacheInvalidation": "useAuthInv",
        "useAccessCacheInvalidation": "useAccessInv",
        "useAuditCacheInvalidation": "useAuditInv",
        "useClientsCacheInvalidation": "useClientsInv",
        "useDirectoryCacheInvalidation": "useDirectoryInv",
        "useImageCacheInvalidation": "useImageInv",
        "useSystemCacheInvalidation": "useSystemInv",
        "useTenantsCacheInvalidation": "useTenantsInv",
        "useCacheInvalidationEvents": "useQueryInv",
    }
    for old, new in mapping.items():
        src = qdir / old
        if src.exists():
            dst = qdir / new
            text = src.read_text()
            for a, b in fn_map.items():
                text = text.replace(a, b)
            dst.write_text(text)
            src.unlink()
            print("inv rename", old, "->", new)

    qp = SRC / "providers/QueryProvider/index.tsx"
    qp.write_text(
        qp.read_text()
        .replace("./hooks/useCacheInvalidationEvents", "./hooks/useQueryInv")
        .replace("useCacheInvalidationEvents", "useQueryInv")
    )


def main() -> None:
    hooks_dir = SRC / "hooks"
    for old, (hook_fn, alias, _new) in HOOK_MAP.items():
        path = hooks_dir / old
        if path.exists():
            convert_hook(path, hook_fn, alias)
    rename_hooks()
    rename_inv()


if __name__ == "__main__":
    main()
