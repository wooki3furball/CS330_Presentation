# CS330_Presentation

Author: Michael Bopp <br>
Last Edited: 8/29/23 <br>

Presentation Link:
--- 
https://docs.google.com/presentation/d/1GOgSLNMm2SAIfpOeIdAmaBRCgapucFF09w2sps8X9SA/edit?usp=sharing

# Downloading Docker

If you are using any version of Windows, specific WSL dependencies must be installed to run Docker software.

Likewise any flavor of Unix: BSD, GNU/Linux, MacOSX has specific instructions on the internet.

I recommend reading or watching a short tutorial to download & start Docker Desktop on your specific OS.

You can proceed with the rest of this tutorial without Docker, you'll just need to download the official Go compiler & Python interpretter.

# Prepare Environment Variables

If you are planning to use Docker you'll need to set an environment variable.

```powershell
# Windows
# Path to project: i.e. <driveletter:\...\CS330_Presentation>
$env:BUILD_CONTEXT="insert file path here"

# Clear env var if a mistake occurs
Remove-Item Env:BUILD_CONTEXT
```

```bash
# Unix OSes
# Path to project
export BUILD_CONTEXT="/path/to/directory"

# Might need to source or restart shell after export
source ~/.bashrc

# Clear env var
unset BUILD_CONTEXT
```

# Interacting with Docker through dscript.py

A python interpretter is needed, your command might be python3 instead of python.

```bash
# Build a Docker image
python dscript.py build

# Build a Docker container & attach to the container in a shell
python dscript.py run
```

# Taskfile execution

By this point you should be in a Docker container, or have Go & Python installed.

![Taskfile](taskfile_demo.jpg)


# Running wasm

```bash
task --list-all

# Compile Go -> wasm code
task b

# Start the web test server
task s

```

If you are in a container manually go to the link below. <br> http://localhost:80

Then you must click the wasm folder.

# Sources

JS code I based my Go adaptation on: https://dev.to/foqc/mandelbrot-set-in-js-480o <br>
https://dev.to/lydiahallie/javascript-visualized-promises-async-await-5gke <br>
Wasm code explorer: https://wasdk.github.io/wasmcodeexplorer/ <br>
More wasm stuff: https://paoloseverini.wordpress.com/
