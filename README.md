# Setup instructions for obsidian Vault with local AI

## Ollama models

Coding
https://ollama.com/library/qwen3.8

General purpose
https://ollama.com/library/gemma4

## Tool instructions

To use Ollama with Obsidian, you need to install Ollama locally, download an AI model, and connect it to Obsidian via a community plugin. This configuration keeps your data entirely private on your machine without subscription fees.

## 1. Step 1: Install Ollama and Download a Model

1. Go to the Ollama Official Website to download and install the app for Mac, Windows, or Linux.
2. Open your terminal (macOS/Linux) or Command Prompt (Windows).
3. Download a lightweight language model by running the following command:

```bash
ollama pull llama3
```

Verify the model downloaded successfully:

``` bash
ollama list
```

## Step 2: Configure Ollama for Obsidian (CORS Settings)

Obsidian requires special permissions (CORS) to securely connect to a local server running on your machine.

- macOS: In your terminal, run:

``` bash
launchctl setenv OLLAMA_ORIGINS "app://obsidian.md*"
```

## Step 3: Install a Plugin in Obsidian

Copilot (for a ChatGPT-like sidebar interface) and Local GPT (for inline text processing).

1. Obsidian and navigate to Settings -> Community plugins.
2. Click Turn on community plugins if you haven't already.
3. Click Browse and search for either Copilot or Local GPT.
4. Click Install, then click Enable.

## Step 4: Link the Plugin to Ollama

Option A: If using the "Copilot" Plugin
1. Go to Settings -> Copilot.
2. Change the API Provider dropdown menu to Ollama (or Self-hosted Ollama API).
3. Ensure the default URL is set to http://localhost:11434.
4. In the Model field, type the exact name of the model you downloaded (e.g., llama3:latest or deepseek-r1).
5. Click Verify/Test Connection.

Option B: If using the "Local GPT" Plugin
1. Go to Settings -> AI Providers (installed automatically with Local GPT).
2. Click + to add a new provider and select Ollama.
3. Click the refresh icon next to the model field to automatically fetch your local Ollama models.
4. Click Save.
5. Go to Settings -> Local GPT and select your Ollama model from the provider dropdowns.

## Step 5: Start Using AI in Your Notes