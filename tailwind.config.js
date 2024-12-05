/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./pkg/ui/**/*.templ"],
    theme: {
        extend: {},
    },
    safelist: [
        "bg-gray-100",
        "flex",
        "h-32",
        "h-4",
        "h-48",
        "items-center",
        "rounded",
        "rounded-full",
        "rounded-lg",
        "space-x-4",
        "space-y-2",
        "space-y-3",
        "space-y-6",
        "w-1/2",
        "w-1/3",
        "w-2/3",
        "w-32",
        "w-4/6",
        "w-5/6",
        "w-full"
    ],
    plugins: [],
}