import { Constants } from '@/constants/constants';
import { RedditToken } from '@/models/RedditToken';
import { v4 as uuid } from 'uuid';
import { NotifierService } from './notifier-service';

export class RedditService {
    // All bearer tokens expire after 1 hour.
    private static BASE_URL = 'https://www.reddit.com/api/v1';
    private static CLIENT_ID = process.env.CLIENT_ID!;
    private static CLIENT_SECRET = process.env.CLIENT_SECRET!;
    private static REDIRECT_URI = process.env.REDIRECT_URI!;

    public static getAuthUrl(): string {
        const newUUID: string = uuid();
        const params = new URLSearchParams({
            client_id: this.CLIENT_ID,
            response_type: 'code',
            state: newUUID,
            redirect_uri: this.REDIRECT_URI,
            duration: 'permanent',
            scope: 'identity read',
        });
        return `${this.BASE_URL}/authorize?${params.toString()}`;
    }

    public static async fetchRedditToken(
        code: string
    ): Promise<RedditToken | null> {
        try {
            const response = await fetch(`${this.BASE_URL}/access_token`, {
                method: 'POST',
                headers: {
                    Authorization: `Basic ${btoa(
                        `${this.CLIENT_ID}:${this.CLIENT_SECRET}`
                    )}`,
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `grant_type=authorization_code&code=${code}&redirect_uri=${this.REDIRECT_URI}`,
            });

            if (response.ok) {
                const redditToken: RedditToken = await response.json();
                redditToken.created_at = Date.now();
                const redditTokenJSONString: string =
                    JSON.stringify(redditToken);
                localStorage.setItem(
                    Constants.REDDIT_TOKEN,
                    redditTokenJSONString
                );
                return redditToken;
            }
            NotifierService.showError('response is not 200!');
            return null;
        } catch {
            NotifierService.showError('Error creating RedditToken');
            return null;
        }
    }

    public static async refreshRedditToken(
        redditToken: RedditToken
    ): Promise<RedditToken | null> {
        try {
            const response = await fetch(`${this.BASE_URL}/access_token`, {
                method: 'POST',
                headers: {
                    Authorization: `Basic ${btoa(
                        `${this.CLIENT_ID}:${this.CLIENT_SECRET}`
                    )}`,
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `grant_type=refresh_token&refresh_token=${redditToken.refresh_token}`,
            });

            if (response.ok) {
                const redditToken: RedditToken = await response.json();
                redditToken.created_at = Date.now();
                localStorage.setItem(
                    Constants.REDDIT_TOKEN,
                    JSON.stringify(redditToken)
                );
                return redditToken;
            }
            NotifierService.showError('response is not 200!');
            return null;
        } catch {
            NotifierService.showError('Error creating RedditToken');
            return null;
        }
    }

    public static async getRedditTokenFromLocalStorage(): Promise<RedditToken | null> {
        const redditTokenJSONString: string | null = localStorage.getItem(
            Constants.REDDIT_TOKEN
        );
        if (redditTokenJSONString) {
            // check if expired
            let redditToken: RedditToken = JSON.parse(redditTokenJSONString);

            if (this.isRedditTokenExpired(redditToken)) {
                // refresh
                const refreshedRedditToken = await this.refreshRedditToken(
                    redditToken
                );
                if (refreshedRedditToken) {
                    redditToken = refreshedRedditToken;
                }
            }

            return redditToken;
        }
        return null;
    }

    public static async getUsername(
        redditToken: RedditToken
    ): Promise<string | null> {
        try {
            const response = await fetch(`https://oauth.reddit.com/api/v1/me`, {
                method: 'GET',
                headers: {
                    Authorization: `Bearer ${redditToken.access_token}`,
                    'User-Agent': 'web_opt_downloader/0.1 by lost_in_trap',
                },
            });

            if (response.ok) {
                const jsonObjects: any = await response.json();
                return jsonObjects.name;
            }
            NotifierService.showError('response is not 200!');
            return null;
        } catch {
            NotifierService.showError('Error getting Username');
            return null;
        }
    }

    public static async getHotPost(
        redditToken: RedditToken,
        subreddit: string
    ) {
        try {
            const response = await fetch(
                `https://oauth.reddit.com/r/${subreddit}/hot?raw_json=1`,
                {
                    method: 'GET',
                    headers: {
                        Authorization: `Bearer ${redditToken.access_token}`,
                        'User-Agent': 'web_opt_downloader/0.1 by lost_in_trap',
                    },
                }
            );

            if (response.ok) {
                const jsonObjects: any = await response.json();
                return jsonObjects.data.children;
            }
            NotifierService.showError('response is not 200!');
        } catch {
            NotifierService.showError('Error getting Post');
        }
    }

    public static deleteTokenFromLocalStorage() {
        localStorage.removeItem(Constants.REDDIT_TOKEN);
    }

    private static isRedditTokenExpired(redditToken: RedditToken): boolean {
        const currentTime = Date.now();
        const expiryTime =
            redditToken.created_at + redditToken.expires_in * 1000; // expires_in is in seconds, convert to milliseconds
        return currentTime >= expiryTime;
    }
}
