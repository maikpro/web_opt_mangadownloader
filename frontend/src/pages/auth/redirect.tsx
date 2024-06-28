import { RedditService } from '@/services/reddit-service';
import { useRouter } from 'next/router';
import { useEffect } from 'react';

export default function RedditAuth() {
    const router = useRouter();
    useEffect(() => {
        const urlParams = new URLSearchParams(window.location.search);
        const code = urlParams.get('code');
        const state = urlParams.get('state');

        if (code && state) {
            const getRedditToken = async () => {
                await RedditService.fetchRedditToken(code);
                router.push('/');
            };
            getRedditToken();
        }
    }, []);

    return (
        <div>
            <h1>Redirected from Reddit...</h1>
        </div>
    );
}
