#!/usr/bin/env python

import os.path
import subprocess

def main():
    dependencies = [
        "https://github.com/hashicorp/hcl.git",
    ]
    pwd = os.getcwd()

    for d in dependencies:
        p = os.path.join("vendor", d.split("/")[2], d.split("/")[3])
        if not os.path.isdir(p):
            os.makedirs(p)
        os.chdir(p)
        cmd = "git clone " + d
        subprocess.call(cmd.strip().split(" "))
        os.chdir(pwd)

if __name__ == '__main__':
    main()

