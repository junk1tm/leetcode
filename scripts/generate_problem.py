#!/usr/bin/env python3

import json
import sys
import unicodedata
from html import parser
from urllib import request


class HTMLFilter(parser.HTMLParser):
    """
    Convert HTML to plain text.
    """

    def __init__(self) -> None:
        self.text = ""
        super().__init__()

    def handle_data(self, data: str) -> None:
        self.text += data

    @classmethod
    def html2text(cls, html: str) -> str:
        f = cls()
        f.feed(html)
        return f.text


GRAPHQL_QUERY = """\
query questionData($titleSlug: String!) {
    question(titleSlug: $titleSlug) {
        questionId
        title
        difficulty
        content
        codeSnippets {
            langSlug
            code
        }
    }
}
"""


def main() -> None:
    title = sys.argv[1]  # in kebab-case

    payload = {
        "operationName": "questionData",
        "variables": {"titleSlug": title},
        "query": GRAPHQL_QUERY,
    }
    body = json.dumps(payload).encode()

    headers = {
        "Content-Type": "application/json",
        "User-Agent": "Mozilla/5.0",
    }

    req = request.Request("https://leetcode.com/graphql", data=body, headers=headers, method="POST")
    with request.urlopen(req) as resp:
        data = json.load(resp)["data"]["question"]

    number = int(data["questionId"])
    html = str(data["content"])
    text = HTMLFilter.html2text(html)
    text = text.replace("\n", "\n// ")  # comment all lines
    text = unicodedata.normalize("NFKD", text)  # remove <0xa0>

    code = "package leetcode\n\n"
    for snippet in list(data["codeSnippets"]):
        if snippet["langSlug"] == "golang":
            code += str(snippet["code"])
            break

    filename = f"{number:04}_{title.replace('-', '_')}.go"
    with open(filename, "w") as f:
        f.write(f"""\
// https://leetcode.com/problems/{title}
//
// {number:04}. {data['title']} [{data['difficulty']}]
//
// {text}

{code}
""")


if __name__ == "__main__":
    main()
