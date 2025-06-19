# Claude Code & Cursor Configuration Examples

This directory contains configuration file examples for using the n8n-workflow-mcp server with the Claude Code & Cursor app.

## 📁 File List

### 🚀 Standalone Mode (Recommended)

#### 1. `claude_desktop_config_standalone.json`
- **Purpose**: Use binary built in Standalone mode (recommended)
- **Advantages**: 
  - Run as a single binary without external files
  - Fast startup time
  - Easy deployment
- **Target**: General users, production environment

#### 2. `claude_desktop_config_standalone_en.json`
- **Purpose**: Run Standalone mode with English interface
- **Features**: All tool descriptions displayed in English
- **Target**: English-speaking users

#### 3. `claude_desktop_config_standalone_jp.json`
- **Purpose**: Run Standalone mode with Japanese interface
- **Features**: All tool descriptions displayed in Japanese
- **Target**: Japanese users

### 🛠 Development Mode

#### 4. `claude_desktop_config_dev_filesystem.json`
- **Purpose**: Run in filesystem mode for development
- **Advantages**: 
  - Immediate reflection of data file changes when source code is modified
  - Convenient debugging during development
- **Disadvantages**: Dependency on external data files
- **Target**: Developers

#### 5. `claude_desktop_config_dev_go.json`
- **Purpose**: Run directly with Go command (for development)
- **Advantages**: 
  - Automatic compilation when source code changes
  - Real-time reflection during development
- **Disadvantages**: Requires Go environment, increased execution time
- **Target**: Developers

### 🪟 Windows Only

#### 6. `claude_desktop_config_windows_standalone.json`
- **Purpose**: Run Standalone mode on Windows
- **Features**: Uses Windows path format
- **Target**: Windows users

#### 7. `claude_desktop_config_windows_dev.json`
- **Purpose**: Run with Go command for development on Windows
- **Features**: Windows path format and development options
- **Target**: Windows developers

## 🎯 Recommended Settings

### General Users
1. **macOS/Linux**: `claude_desktop_config_standalone.json`
2. **Windows**: `claude_desktop_config_windows_standalone.json`

### Developers
1. **During development**: `claude_desktop_config_dev_go.json`
2. **During testing**: `claude_desktop_config_dev_filesystem.json`

### Multilingual Users
1. **English**: `claude_desktop_config_standalone_en.json`
2. **Japanese**: `claude_desktop_config_standalone_jp.json`
3. **Korean**: `claude_desktop_config_standalone.json` (default)

## 🛠️ Available Tools

### 1. `search_n8n_nodes`
- **Purpose**: n8n node search
- **Function**: Search for n8n nodes based on keywords and return results with scores
- **Parameters**:
  - `query` (required): Keywords to search for
  - `limit` (optional): Maximum number of results to return (default: 5, max: 20)

### 2. `search_workflow` ✨ **Newly Added!**
- **Purpose**: n8n workflow template search
- **Function**: Search community workflow templates and provide download URLs
- **Parameters**:
  - `query` (required): Keywords to search for (e.g., "AI", "automation", "chatbot")
  - `max_results` (optional): Maximum number of results to return (default: 5, max: 15)
  - `category` (optional): Category filter (e.g., airtable, blog, common, crm, email)
- **Example Usage**:
  - Find AI chatbot workflows: `query: "AI chatbot"`
  - Email automation workflows: `query: "email automation", category: "email"`
  - Airtable integration workflows: `category: "airtable"`

### 3. `validate_workflow`
- **Purpose**: n8n workflow JSON validation
- **Function**: Validate syntax and structure of workflow JSON files
- **Parameters**:
  - `file_path` (required): Path to JSON file to validate
  - `include_line_numbers` (optional): Include detailed line-by-line information (default: true)
  - `auto_fix_suggestions` (optional): Provide automatic fix suggestions (default: true)

### 4. `ping`
- **Purpose**: Server status check
- **Function**: Basic test to verify MCP server is working properly

## 🔧 Configuration Method

### 1. Check Configuration File Location
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

### 2. Select and Apply Configuration File
1. Select the appropriate configuration file from the file list above.
2. Copy the contents of the selected file.
3. Paste it into the Claude Code & Cursor configuration file.
4. Modify the file paths to match your environment.

### 3. Path Modification Guide

**macOS/Linux Example:**
```json
"command": "/Users/john/projects/n8n-workflow-mcp/build/n8n-mcp-server"
```

**Windows Example:**
```json
"command": "C:\\Users\\john\\projects\\n8n-workflow-mcp\\build\\n8n-mcp-server.exe"
```

### 4. Restart Claude Code & Cursor
After applying the configuration, completely close the Claude Code & Cursor app and restart it.

## ⚙️ Advanced Configuration Options

### Search Weight Adjustment
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "/path/to/n8n-mcp-server",
      "args": [
        "--standalone",
        "--lang", "en",
        "--name-weight", "3.0",
        "--overview-weight", "2.0",
        "--case-weight", "1.5",
        "--op-weight", "1.0"
      ]
    }
  }
}
```

### Custom Data Path (Filesystem Mode)
```json
{
  "mcpServers": {
    "n8n-kor-workflow": {
      "command": "/path/to/n8n-mcp-server",
      "args": [
        "-standalone=false",
        "--data-path", "/custom/path/to/nodes",
        "--lang", "en"
      ]
    }
  }
}
```

## 🔍 Troubleshooting

### Common Issues

1. **Server won't start**
   - Check if file path is correct
   - Check if execution permission exists (`chmod +x`)
   - Check if built binary exists

2. **Permission errors occur**
   ```bash
   chmod +x /path/to/n8n-mcp-server
   ```

3. **Path-related errors**
   - Recommend using absolute paths
   - On Windows, backslashes need to be escaped: `\\`

### Check Logs
To check Claude Code & Cursor logs:
- **macOS**: Search for "Claude" in Console.app
- **Windows**: Check Event Viewer or Claude log files

## 📝 Notes

- All paths must use absolute paths
- On Windows, when using backslashes (`\`) in paths, enter them twice: `\\`
- Or you can use forward slashes (`/`)
- Always restart Claude Code & Cursor after configuration changes
- Development mode requires Go environment to be installed

## 🆘 Help

If you need additional help:
1. Refer to the main README.md of the project
2. Contact via GitHub Issues
3. Configuration file syntax validation: [JSON Validator](https://jsonlint.com/) 