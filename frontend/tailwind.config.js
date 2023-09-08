/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./index.html","./src/**/*.{html,js}","./src/**/*.{vue,js,ts,jsx,tsx,svelte}",],
    theme: {
      extend: {},
    },
    plugins: [require("daisyui")],
    daisyui: {
      themes: [
        {
          mytheme: {
            "primary": "#2563eb",
            "secondary": "#f97316",
            "accent": "#1fb2a6",
            "neutral": "#2a323c",
            "base-100": "#1d232a",
            "info": "#3abff8",
            "success": "#36d399",
            "warning": "#fbbd23",
            "error": "#f87272",
          },
        },
      ],
    },
  }