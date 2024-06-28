/** @type {import('next').NextConfig} */
const nextConfig = {
    env: {
        CLIENT_ID: process.env.CLIENT_ID,
        CLIENT_SECRET: process.env.CLIENT_SECRET,
        REDIRECT_URI: process.env.REDIRECT_URI,
    },
};

export default nextConfig;
