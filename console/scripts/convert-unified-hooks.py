#!/usr/bin/env python3
"""Convert makeUnifiedQuery factory calls to CityPulso useUnifiedQuery."""
from __future__ import annotations

import pathlib

ROOTS = [
    pathlib.Path("/volume1/Projects/NebulaForgeX/NFX-Identity/console/src/hooks"),
    pathlib.Path("/volume1/Projects/NebulaForgeX/NFX-Vault/frontend/src/hooks"),
    pathlib.Path("/volume1/Projects/NebulaForgeX/NFX-Storages/storage-console/src/hooks"),
]


def split_args(inner: str) -> list[str]:
    args: list[str] = []
    buf: list[str] = []
    depth = 0
    in_str = None
    escape = False
    for ch in inner:
        if in_str:
            buf.append(ch)
            if escape:
                escape = False
            elif ch == "\\":
                escape = True
            elif ch == in_str:
                in_str = None
            continue
        if ch in "\"'`":
            in_str = ch
            buf.append(ch)
            continue
        if ch in "([{":
            depth += 1
            buf.append(ch)
            continue
        if ch in ")]}":
            depth -= 1
            buf.append(ch)
            continue
        if ch == "," and depth == 0:
            args.append("".join(buf).strip())
            buf = []
            continue
        buf.append(ch)
    if buf:
        args.append("".join(buf).strip())
    return args


def match_parens(text: str, open_idx: int) -> int | None:
    depth = 0
    in_str = None
    escape = False
    for i, ch in enumerate(text[open_idx:], open_idx):
        if in_str:
            if escape:
                escape = False
            elif ch == "\\":
                escape = True
            elif ch == in_str:
                in_str = None
            continue
        if ch in "\"'`":
            in_str = ch
            continue
        if ch == "(":
            depth += 1
        elif ch == ")":
            depth -= 1
            if depth == 0:
                return i
    return None


def convert_file(path: pathlib.Path) -> bool:
    text = path.read_text()
    original = text
    changed = False

    while True:
        inf = False
        idx = text.find("const makeQuery = makeUnifiedQuery(")
        if idx < 0:
            idx = text.find("const makeQuery = makeUnifiedInfiniteQuery")
            if idx < 0:
                break
            inf = True
        start = text.find("(", idx)
        end = match_parens(text, start)
        if end is None:
            break
        inner = text[start + 1 : end]
        rest = text[end + 1 :]
        leading = len(rest) - len(rest.lstrip())
        after = rest.lstrip()
        if after.startswith(";"):
            after = after[1:].lstrip()
            leading += 1
            leading += len(rest.lstrip()[1:]) - len(rest.lstrip()[1:].lstrip()) if False else 0
        if not after.startswith("return makeQuery("):
            print("skip unmatched", path)
            break
        call_start = after.find("(")
        call_end = match_parens(after, call_start)
        if call_end is None:
            break
        call_inner = after[call_start + 1 : call_end]
        factory_args = split_args(inner)
        call_args = split_args(call_inner)
        fetch_fn = factory_args[0]
        mode = factory_args[1].strip().strip("\"'") if len(factory_args) > 1 else "normal"
        extra = factory_args[2] if len(factory_args) > 2 else None
        key = call_args[0] if call_args else "[]"
        filt = call_args[1] if len(call_args) > 1 else "{}"
        options = call_args[2] if len(call_args) > 2 else "undefined"

        if inf:
            hook = "useUnifiedSuspenseInfiniteQuery" if mode == "suspense" else "useUnifiedInfiniteQuery"
            replacement = f"return {hook}(\n    {fetch_fn},\n    {key},\n    {filt},\n    {options},\n  )"
        else:
            hook = "useUnifiedSuspenseQuery" if mode == "suspense" else "useUnifiedQuery"
            if extra and extra not in ("undefined",):
                replacement = (
                    f"return {hook}(\n    {fetch_fn},\n    {key},\n    {filt},\n    {{ ...({options} ?? {{}}), postProcess: {extra} }},\n  )"
                )
            else:
                replacement = f"return {hook}(\n    {fetch_fn},\n    {key},\n    {filt},\n    {options},\n  )"

        rest2 = text[end + 1 :]
        lead = len(rest2) - len(rest2.lstrip())
        after2 = rest2.lstrip()
        if after2.startswith(";"):
            after2 = after2[1:]
            lead += 1
            ws = len(after2) - len(after2.lstrip())
            lead += ws
            after2 = after2.lstrip()
        call_start2 = after2.find("(")
        call_end2 = match_parens(after2, call_start2)
        if call_end2 is None:
            break
        return_end = end + 1 + lead + call_end2 + 1
        if after2[call_end2 + 1 :].startswith(";"):
            return_end += 1
        text = text[:idx] + replacement + text[return_end:]
        changed = True

    if "makeUnifiedQuery" in text:
        text = text.replace("makeUnifiedQuery", "useUnifiedQuery")
        changed = True
    if "makeUnifiedInfiniteQuery" in text:
        text = text.replace("makeUnifiedInfiniteQuery", "useUnifiedInfiniteQuery")
        changed = True
    if "from \"nfx-ui/hooks\"" in text and "useUnifiedQuery" in text and "useUnifiedSuspenseQuery" not in text:
        text = text.replace("useUnifiedQuery", "useUnifiedQuery, useUnifiedSuspenseQuery")
        changed = True

    if changed and text != original:
        path.write_text(text)
        print("converted", path)
        return True
    return False


def main() -> None:
    for root in ROOTS:
        if not root.exists():
            continue
        for path in list(root.rglob("*.ts")) + list(root.rglob("*.tsx")):
            convert_file(path)


if __name__ == "__main__":
    main()
