# Websys
A web-based app suite with a Go backend for interactting with the underlying system.

## Description

Websys presents a suite of apps similar to those that you use every day (text
editor, image viewer, etc.), with one difference — these apps run in the
browser. They can still interact with your underlying operating system, but they
are entirely web-based.

## Setup

```
git clone https://git.sr.ht/~kvo/websys
cd websys/src/
go build -o ../prg/websys .
```

Then you can execute `../prg/websys` to start the Websys backend.

## Usage

When you start the Websys backend, you will be prompted for the username and
password with which you wish to sign into Websys. Websys needs this information
so that if you are logging into Websys from a remote computer, your connections
are authenticated. The username and password don't need to match your system
credentials, but please ensure they are strong to minimise any chance of system
compromise.

Open <http://localhost:2038>, sign in with the same credentials as the ones you
gave Websys initially, and enjoy!

## Contributing

The backend server source is stored in `src/`. Each subfolder in `src/` (except
for `src/auth/`) holds the code for a separate app. This includes:

  - the server handler for the app (`src/<APP>/main.go`)
  - the HTML for the app window (`src/<APP>/ui.html`)
  - the CSS for the app layout and theming (`src/<APP>/styles.css`)
  - the JS for the app (`src/<APP>/main.js`)

Other auxiliary HTML/CSS/JS/Go files may be present if they are necessary.

Please send patches to <~kvo/websys@lists.sr.ht>
