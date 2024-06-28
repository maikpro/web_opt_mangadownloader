import type { AppProps } from 'next/app';
import '../public/globals.css';
import '../public/images/favicon.ico';
import Layout from '@/components/layout';

export default function MyApp({ Component, pageProps }: AppProps) {
    return (
        <Layout>
            <Component {...pageProps} />
        </Layout>
    );
}
