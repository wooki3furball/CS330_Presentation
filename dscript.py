# Author: Michael Bopp
# Filename: dscript.py
# Description: Python program to configure Dockerfile with build & run commands.
# Date Created: 8/29/23
# Last Edited: 8/29/23
# Dependency: Python Interpreter

import os
import subprocess
import sys
import argparse

def docker_build(no_cache=False):
    """
    Construct the Docker build command
    """
    docker_build_command = [
        'docker', 'build',
        '-t', 'alpine:CS330',
        '.'  # Build context (current directory)
    ]
    
    if no_cache:
        docker_build_command.append('--no-cache')
    
    subprocess.run(docker_build_command)

def docker_run(env_var, port_mapping):
    """
    Construct the Docker run command
    """
    docker_run_command = [
        'docker', 'run',
        '-v', f'{env_var}:/workspace', # Mount storage volume from host OS to container
        '-p', f'{port_mapping}:80',
        '-it',  # Interactive terminal
        'alpine:CS330'
    ]
    
    subprocess.run(docker_run_command)

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("action", choices=["build", "run"], help="Either 'build' or 'run'")
    parser.add_argument("--no-cache", action="store_true", help="Use --no-cache with Docker build")
    args = parser.parse_args()
    
    if args.action == 'build':
        docker_build(no_cache=args.no_cache)
    elif args.action == 'run':
        env_var = os.environ.get("BUILD_CONTEXT")
        if env_var is None:
            sys.exit("There was no environment variable.")
        port_mapping = '8080'
        docker_run(env_var, port_mapping)