/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,ts}",
  ],
  theme: {
    extend: {
      colors: {
        'main' : '#464655',
        'redish': '#FF4B55',
        'yellowish': '#FFE15A',
        'accent': '#8EA4D2',
        'whiteish': '#CEE7E6'
      },
      fontFamily: {
        'para': ['Taviraj'],
        'small-header': ['Fenix'],
        'header': ['Prata'] 
      }
    },
  },
  plugins: [],
}