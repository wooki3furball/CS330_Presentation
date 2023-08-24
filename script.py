# Author: Michael Bopp
# Filename: script.py
# Description: Python program to be configured in Dockerfile.
# Date Created: 8/23/23
# Last Edited: 8/23/23
# Dependency: Python Interpreter

import subprocess

def execute_commands():
    """
    Download go-task and alias it with task via a symbolic link.
    """
    try:
        subprocess.run(['apk', 'add', 'go-task'], check=True)
        go_task_path = subprocess.getoutput('which go-task')
        subprocess.run(['ln', '-s', go_task_path, '/usr/bin/task'], check=True)
        print("Commands executed successfully!")
    except subprocess.CalledProcessError as e:
        print(f"An error occurred while executing the commands: {e}")

if __name__ == "__main__":
    execute_commands()