/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./dist/index.html","../components/*.go","./src/**/*.{vue,html,js,ts,jsx,tsx,svelte}",],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    logs: false,
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
