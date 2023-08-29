# Author: Michael Bopp
# Filename: dscript.py
# Description: Python program to configure Dockerfile with build & run commands.
# Date Created: 8/29/23
# Last Edited: 8/29/23
# Dependency: Python Interpreter

import os
import subprocess
import sys

def docker_build():
    """
    Construct the Docker build command
    """
    docker_build_command = [
        'docker', 'build',
        '-t', 'alpine:CS330',
        '.'  # Build context (current directory)
        # Need to add the --no-cache command to the build, also need to make it possible to pass into the script at CLI
    ]
    
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
    if len(sys.argv) != 2:
        print("Usage: python dscript.py [build|run]")
        sys.exit(1)
    
    action = sys.argv[1].lower()
    
    if action == 'build':
        docker_build()
    elif action == 'run':
        # Retrieve the environment variable and port number
        env_var = os.environ["BUILD_CONTEXT"]
        if env_var == None:
            sys.exit("There was no environment variable.")
        port_mapping = '8080'
        docker_run(env_var, port_mapping)
    else:
        print("Invalid command. Use either 'build' or 'run'.")