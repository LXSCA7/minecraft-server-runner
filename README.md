# Minecraft Server Runner

A Go script to connect to a Virtual Machine and manage a Minecraft server.

## Features
- SSH connection to your VM.
- Start and stop a Minecraft server using a CLI menu.
- Status monitoring for the server.

## Requirements
- SSH key configured to access the target VM.
- A Minecraft server set up on your VM.

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/LXSCA7/minecraft-server-runner.git
   cd minecraft-server-runner
   ```
2. Create a settings.json file in the project root, based on the template at `settings.json.example`
```json
{
    "vm_user": "your_vm_username",
    "vm_ip": "your_vm_ip_address",
    "key_path": "path_to_your_ssh_key",
    "server_path": "path_to_your_minecraft_server",
    "java_command": "java -Xmx16G -Xms16G -jar server.jar nogui"
}
```
3. Run.
```bash
go run .
```

## Notes
- Ensure your settings.json file is correctly configured before running the script.
- The script uses the screen command to manage the server. Ensure screen is installed on your VM.
