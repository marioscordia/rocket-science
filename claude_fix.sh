SHELL_CONF=""
case "$SHELL" in
    */zsh)  SHELL_CONF="$HOME/.zshrc" ;;
    */bash) SHELL_CONF="$HOME/.bashrc" ;;
    *)      echo "Unknown shell. Please manually add to your config." ;;
esac

if [ -n "$SHELL_CONF" ]; then
    if ! grep -q ".local/bin" "$SHELL_CONF"; then
        echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$SHELL_CONF"
        echo "Added to $SHELL_CONF"
    else
        echo "Path already exists in $SHELL_CONF"
    fi
    source "$SHELL_CONF"
    echo "Shell reloaded. Try running 'claude' now."
fi
