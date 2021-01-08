# Zonelet HTTP Server

### Simplest static site HTTP server in Go, including the instructions on how to set it up, and publish the content.

### Think a __cheatsheet__ of setting up a simple HTTP server self-hosted.

# INSTRUCTIONS:

### 1. Build main.go to Zonelet executable

With this command:

    $ go build

You must have had Go language installed.

### 2. Put Zonelet executable on a server

    $ ssh root@myserver.com
    # mkdir /data/mysrv1
    # exit
    
    $ rsync -r --delete . root@myserver.com:/data/mysrv1/

### 3. Install the server as a systemd service on a server

    $ ssh root@myserver.com
    # cd /data/mysrv1
    # cp ZoneletHTTP.service /etc/systemd/system/
    # nano /etc/systemd/system/

Edit the file, specifically edit the address, as wanted:

    ExecStart=/data/http-serve-1/Zonelet -addr=":49999"

Save it, and continue.

    # systemctl enable ZoneletHTTP
    # systemctl restart ZoneletHTTP

Make sure it is running:

    # systemctl status ZoneletHTTP
    # journalctl -u ZoneletHTTP

    # exit

### 4. Publish static HTML content on a server

Go to a dir with your HTML pages, then:

    $ rsync -r --delete . root@myserver.com:/data/mysrv1/static/

### 5. Done.
