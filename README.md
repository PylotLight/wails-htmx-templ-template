# README

## About

This template uses a unique combination of pure htmx for interactivity plus Go templates for creating components and forms, also included:
- Built-in added styling to show off some of tailwind and daisyui.
- Built-in version display linked to version variable from main which can be updated on app build and during runtime via a function.


## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `bun run dev`. The frontend dev server will run on http://localhost:34115. Connect to this in your
browser and connect to your application.

## Building

To build a redistributable, production mode package, use `wails build`.
