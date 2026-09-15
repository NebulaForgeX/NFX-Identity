#!/usr/bin/env python3
from pathlib import Path
import re

HOOKS = {
    "auth.ts": "auth",
    "directory.ts": "directory",
    "access.ts": "access",
    "audit.ts": "audit",
    "clients.ts": "clients",
    "image.ts": "image",
    "system.ts": "system",
    "tenants.ts": "tenants",
}
ROOT = Path("/volume1/Projects/NebulaForgeX/NFX-Identity/console/src/hooks")
HOOK_FN = {
    "auth": "useAuthRepository",
    "directory": "useDirectoryRepository",
    "access": "useAccessRepository",
    "audit": "useAuditRepository",
    "clients": "useClientsRepository",
    "image": "useImageRepository",
    "system": "useSystemRepository",
    "tenants": "useTenantsRepository",
}

SKIP_PREFIX = re.compile(
    r"^(auth|directory|access|audit|clients|image|system|tenants|console|window|JSON|Error|Promise|Object|Array|Math|Date|Number|String|Boolean)\."
)

for filename, alias in HOOKS.items():
    path = ROOT / filename
    text = path.read_text()
    # prefix leftover PascalCase API calls
    def prefix(m: re.Match[str]) -> str:
        name = m.group(1)
        if name.endswith("Request") or name.endswith("Response") or name.endswith("Params"):
            return m.group(0)
        if name.startswith("use"):
            return m.group(0)
        return f"await {alias}.{name}("

    text = re.sub(r"await (?!(?:auth|directory|access|audit|clients|image|system|tenants)\.)([A-Z]\w+)\(", prefix, text)

    # insert repo in functions that use alias. but don't declare it
    fn_pat = re.compile(r"export const use\w+ = \((?:[^)]|\([^)]*\))*?\) => \{", re.S)

    def insert(m: re.Match[str]) -> str:
        header = m.group(0)
        return header + f"\n  const {alias} = {HOOK_FN[alias]}();"

    # only add if next ~800 chars use alias. but not already declared
    pieces = []
    last = 0
    for m in re.finditer(r"export const use\w+ = ", text):
        start = m.start()
        pieces.append(text[last:start])
        # find matching function body start
        brace = text.find("{", m.end() - 1)
        header = text[start : brace + 1]
        # find end of function roughly next export or EOF
        nxt = text.find("\nexport const use", brace)
        body = text[brace + 1 : nxt if nxt > 0 else len(text)]
        last = brace + 1
        if f"const {alias} =" not in header + body[:400] and f"{alias}." in body:
            header = header + f"\n  const {alias} = {HOOK_FN[alias]}();"
        pieces.append(header)
    pieces.append(text[last:])
    text = "".join(pieces)

    # drop unused useUnifiedQuery if not used as a call
    if "useUnifiedQuery(" not in text:
        text = text.replace("useUnifiedQuery, ", "").replace(", useUnifiedQuery", "")
    if "useTranslation(" not in text:
        text = text.replace('import { useTranslation } from "react-i18next";\n', "")
    path.write_text(text)
    print("fixed", filename)
