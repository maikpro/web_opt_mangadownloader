export interface RedditToken {
    access_token: string;
    expires_in: number; // 86400 seconds => 24 hours
    refresh_token: string;
    scope: string;
    token_type: string;
    created_at: number;
}
