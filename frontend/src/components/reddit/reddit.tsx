import { RedditToken } from '@/models/RedditToken';
import { RedditService } from '@/services/reddit-service';
import { useEffect, useState } from 'react';
import RedditIcon from '../reddit-icon/reddit-icon';
import Link from 'next/link';
import { Url } from 'next/dist/shared/lib/router/router';
import React from 'react';
import LoadingSpinner from '../loading-spinner/loading-spinner';

export default function Reddit() {
    const [redditAuthUrl, setRedditAuthUrl] = useState<Url>('');
    const [redditToken, setRedditToken] = useState<RedditToken>();
    const [username, setUsername] = useState<string>('');
    const [posts, setPosts] = useState<any[]>([]);

    const checkForRedditToken = async () => {
        const redditToken =
            await RedditService.getRedditTokenFromLocalStorage();
        if (redditToken) {
            setRedditToken(redditToken);
        }
        return redditToken;
    };

    const getUsername = async (redditToken: RedditToken) => {
        const username = await RedditService.getUsername(redditToken);
        if (username) {
            setUsername(username);
        }
    };

    const createdDateInGerman = (dateString: number) => {
        const timestamp = dateString * 1000;
        const date = new Date(timestamp);

        // Format the date to German locale
        return new Intl.DateTimeFormat('de-DE', {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            second: 'numeric',
        }).format(date);
    };

    const parseText = (text: string) => {
        const lines = text.split('\n');
        const linkPattern = /\[([^\]]+)\]\(([^)]+)\)/g;
        const statusPatternCross = /\[\]\/\(cross\)/g;
        const statusPatternTick = /\[\]\/\(tick\)/g;
        const boldPattern = /\*\*(.*?)\*\*/g;
        const headerPattern = /^##(.*)/;

        return lines
            .filter((line) => !line.includes(':-----|:-----|:-----'))
            .map((line, index) => {
                if (line === '---' || line === '-') {
                    return <br key={index} />;
                }

                if (headerPattern.test(line)) {
                    const headerText = line.replace(headerPattern, '$1');
                    return (
                        <div key={index} className="text-xl">
                            {headerText}
                        </div>
                    );
                }

                let parsedLine = line.replace(
                    linkPattern,
                    (_, text, url) =>
                        `<a target="_blank" class="decoration-solid text-orange-400" href="${url}">${text}</a>`
                );
                parsedLine = parsedLine.replace(statusPatternCross, '❌');
                parsedLine = parsedLine.replace(statusPatternTick, '✅');
                parsedLine = parsedLine.replace(
                    boldPattern,
                    '<strong>$1</strong>'
                );

                return (
                    <div
                        key={index}
                        dangerouslySetInnerHTML={{ __html: parsedLine }}
                    />
                );
            });
    };

    useEffect(() => {
        setRedditAuthUrl(RedditService.getAuthUrl());
        const initialize = async () => {
            const redditToken = await checkForRedditToken();
            if (redditToken) {
                await getUsername(redditToken);
                const posts = await RedditService.getHotPost(
                    redditToken,
                    'OnePiece'
                );
                setPosts(posts);
            }
        };
        initialize();
    }, []);
    return (
        <div className="mb-10">
            {redditToken ? (
                <div className="rounded-lg border-2 border-amber-500 p-2 overflow-auto h-96">
                    <p className="text-xs mb-5">
                        You are logged into reddit with{' '}
                        <b className="text-orange-500">{username}</b>
                        <button
                            className="ml-2 text-red-700"
                            onClick={() => {
                                RedditService.deleteTokenFromLocalStorage();
                                window.location.href = '/';
                            }}
                        >
                            (Logout)
                        </button>
                    </p>

                    {posts.length > 0 ? (
                        <div>
                            <h1 className="text-xl">Reddit Hot Posts:</h1>
                            {posts.slice(0, 5).map((post, index) => (
                                <div className="mb-5" key={index}>
                                    <a
                                        target="_blank"
                                        className="decoration-solid text-orange-400"
                                        href={post.data.url}
                                    >
                                        <div className="flex text-s items-center">
                                            <RedditIcon />
                                            <p>{post.data.author}</p>
                                        </div>
                                        <div>
                                            {createdDateInGerman(
                                                post.data.created_utc
                                            )}
                                        </div>
                                    </a>
                                    <h1 className="text-xl mt-5 mb-5">
                                        {post.data.title}:
                                    </h1>
                                    <p>{parseText(post.data.selftext)}</p>
                                    <div className="flex mt-5">
                                        <a
                                            target="_blank"
                                            className="decoration-solid text-orange-400"
                                            href={post.data.url}
                                        >
                                            <div className="flex items-center px-2 py-1 rounded-full bg-slate-800">
                                                <svg
                                                    className="mr-1"
                                                    fill="currentColor"
                                                    height="16"
                                                    icon-name="upvote-outline"
                                                    viewBox="0 0 20 20"
                                                    width="16"
                                                    xmlns="http://www.w3.org/2000/svg"
                                                >
                                                    {' '}
                                                    <path d="M12.877 19H7.123A1.125 1.125 0 0 1 6 17.877V11H2.126a1.114 1.114 0 0 1-1.007-.7 1.249 1.249 0 0 1 .171-1.343L9.166.368a1.128 1.128 0 0 1 1.668.004l7.872 8.581a1.25 1.25 0 0 1 .176 1.348 1.113 1.113 0 0 1-1.005.7H14v6.877A1.125 1.125 0 0 1 12.877 19ZM7.25 17.75h5.5v-8h4.934L10 1.31 2.258 9.75H7.25v8ZM2.227 9.784l-.012.016c.01-.006.014-.01.012-.016Z"></path>
                                                </svg>

                                                {post.data.ups}

                                                <svg
                                                    className="ml-1"
                                                    fill="currentColor"
                                                    height="16"
                                                    icon-name="downvote-outline"
                                                    viewBox="0 0 20 20"
                                                    width="16"
                                                    xmlns="http://www.w3.org/2000/svg"
                                                >
                                                    <path d="M10 20a1.122 1.122 0 0 1-.834-.372l-7.872-8.581A1.251 1.251 0 0 1 1.118 9.7 1.114 1.114 0 0 1 2.123 9H6V2.123A1.125 1.125 0 0 1 7.123 1h5.754A1.125 1.125 0 0 1 14 2.123V9h3.874a1.114 1.114 0 0 1 1.007.7 1.25 1.25 0 0 1-.171 1.345l-7.876 8.589A1.128 1.128 0 0 1 10 20Zm-7.684-9.75L10 18.69l7.741-8.44H12.75v-8h-5.5v8H2.316Zm15.469-.05c-.01 0-.014.007-.012.013l.012-.013Z"></path>
                                                </svg>
                                            </div>
                                        </a>

                                        <a
                                            target="_blank"
                                            className="decoration-solid text-orange-400"
                                            href={post.data.url}
                                        >
                                            <div className="flex items-center px-2 py-1 rounded-full bg-slate-800 ml-3">
                                                <svg
                                                    className="mr-1"
                                                    aria-hidden="true"
                                                    fill="currentColor"
                                                    height="20"
                                                    icon-name="comment-outline"
                                                    viewBox="0 0 20 20"
                                                    width="20"
                                                    xmlns="http://www.w3.org/2000/svg"
                                                >
                                                    <path d="M7.725 19.872a.718.718 0 0 1-.607-.328.725.725 0 0 1-.118-.397V16H3.625A2.63 2.63 0 0 1 1 13.375v-9.75A2.629 2.629 0 0 1 3.625 1h12.75A2.63 2.63 0 0 1 19 3.625v9.75A2.63 2.63 0 0 1 16.375 16h-4.161l-4 3.681a.725.725 0 0 1-.489.191ZM3.625 2.25A1.377 1.377 0 0 0 2.25 3.625v9.75a1.377 1.377 0 0 0 1.375 1.375h4a.625.625 0 0 1 .625.625v2.575l3.3-3.035a.628.628 0 0 1 .424-.165h4.4a1.377 1.377 0 0 0 1.375-1.375v-9.75a1.377 1.377 0 0 0-1.374-1.375H3.625Z"></path>
                                                </svg>
                                                {post.data.num_comments}
                                            </div>
                                        </a>
                                    </div>
                                </div>
                            ))}
                        </div>
                    ) : (
                        <LoadingSpinner loadingText="Loading posts from reddit" />
                    )}
                </div>
            ) : (
                <div>
                    <Link href={redditAuthUrl}>
                        <button className="bg-orange-500 hover:bg-orange-600 text-white font-bold py-2 px-4 rounded flex ">
                            <div className="flex items-center">
                                <RedditIcon />
                                <span className="ml-1">Login</span>
                            </div>
                        </button>
                    </Link>
                </div>
            )}
        </div>
    );
}
