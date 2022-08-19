import os
import re
from dataclasses import dataclass


@dataclass
class Problem:
    link: str
    number: int
    title: str
    difficulty: str
    file: str
    test: str = ""


TEMPLATE = """\
# LeetCode

[![ci-img]][ci]
[![license-img]][license]

[ci]: https://github.com/junk1tm/leetcode/actions/workflows/go.yml
[ci-img]: https://github.com/junk1tm/leetcode/actions/workflows/go.yml/badge.svg
[license]: https://github.com/junk1tm/leetcode/blob/main/LICENSE
[license-img]: https://img.shields.io/github/license/junk1tm/leetcode

> A collection of my [LeetCode](https://leetcode.com) solutions written in Go

## Progress

### Total

| Easy | Medium | Hard |
|:----:|:------:|:----:|
|{Easy}|{Medium}|{Hard}|

### Solutions

|  #  | Title | Difficulty | Solution | Tests |
|:---:|:------|:----------:|:--------:|:-----:|
"""


def main():
    problems: list[Problem] = []
    total: dict[str, int] = {"Easy": 0, "Medium": 0, "Hard": 0}

    for filename in sorted(os.listdir()):
        if not re.fullmatch(r"^[0-9]{4}[a-z_]+.go$", filename):
            continue

        with open(filename, "r") as f:
            COMMENT_PREFIX = "// "
            link = f.readline().removeprefix(COMMENT_PREFIX).removesuffix("\n")
            f.readline()
            header = f.readline().removeprefix(COMMENT_PREFIX).removesuffix("\n")

        # the header consists of the task's number, name and difficulty.
        pattern = r"^([0-9]{4})\. ([a-zA-Z ()]+) \[(Easy|Medium|Hard)\]$"
        match = re.fullmatch(pattern, header)
        assert match is not None

        number = int(match.group(1))
        title = match.group(2)
        difficulty = match.group(3)
        file = f"[👀]({filename})"

        p = Problem(link, number, title, difficulty, file)
        problems.append(p)
        total[p.difficulty] += 1

    with open("solutions_test.go", "r") as f:
        it = iter(problems)
        for i, line in enumerate(f):
            if re.match(r"^// [0-9]{4}\.", line):
                next(it).test = f"[✅](solutions_test.go#L{i + 1})"

    with open("README.md", "w") as f:
        f.write(TEMPLATE.format(**total))
        for p in problems:
            f.write(f"|{p.number}|[{p.title}]({p.link})|{p.difficulty}|{p.file}|{p.test}|\n")  # fmt: skip


if __name__ == "__main__":
    main()
