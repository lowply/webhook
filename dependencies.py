#!/usr/bin/env python

import os.path
import subprocess

def main():
    dependencies = [
        "github.com/hashicorp/hcl/1c284ec98f4b398443cbabb0d9197f7f4cc0077c",
        "github.com/mitchellh/cli/5c87c51cedf76a1737bf5ca3979e8644871598a6",
        "github.com/golang/glog/23def4e6c14b4da8ac2ed8007337bc5eb5007998",
    ]
    vendor_dir = "vendor"
    pwd = os.getcwd()
    name = pwd.split("/")[-1]
    git_branch_prefix = "go_vendoring_"

    if not os.path.isdir(vendor_dir):
        os.mkdir(vendor_dir)
        print("created " + vendor_dir + " directory")

    for d in dependencies:
        path = d.split("/")
        repo_name = os.path.join(path[0], path[1], path[2])
        p_creator = os.path.join(vendor_dir, path[0], path[1])
        p_repo    = os.path.join(vendor_dir, path[0], path[1], path[2])

        if not os.path.isdir(p_creator):
            os.makedirs(p_creator)
            print("created " + p_creator + " directory")

        os.chdir(p_creator)

        if not os.path.isdir(path[2]):
            cmd_clone = "git clone https://" + repo_name + ".git"
            subprocess.call(cmd_clone.strip().split(" "))
            print("cloned " + repo_name)

            # check out to target commit with a new branch
            os.chdir(path[2])
            cmd_checkout = "git checkout -b " + git_branch_prefix + name + " " + path[3]
            subprocess.call(cmd_checkout.strip().split(" "))
            print("checked out to " + path[3] + " with new branch " + git_branch_prefix + name)
        else:
            print(p_repo + " is already cloned")

        os.chdir(pwd)

if __name__ == '__main__':
    main()

