import { ChangeEvent, FormEvent, useEffect, useState } from 'react';
import type { Settings } from '@/models/Settings';

export default function Settings() {
    const [settingsData, setSettingsData] = useState<Settings>({
        id: '',
        telegramToken: '',
        telegramChatId: '',
    });

    useEffect(() => {
        const fetchSettings = async () => {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/settings`
            );
            const settings: Settings = await response.json();
            if (settings) {
                setSettingsData(settings);
            }
        };
        fetchSettings();
    }, []);

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setSettingsData({
            ...settingsData,
            [name]: value,
        });
    };

    const save = async (e: FormEvent) => {
        e.preventDefault();
        console.log('Save changes...', settingsData);
        try {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/settings`,
                {
                    method: 'POST',
                    body: JSON.stringify(settingsData),
                }
            );

            if (response.ok) {
                console.log('Data sent successfully');
                const savedSettings: Settings = await response.json();
                setSettingsData(savedSettings);
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };

    const update = async (e: FormEvent) => {
        e.preventDefault();
        console.log('Update changes...', settingsData);
        try {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/settings/id/${settingsData.id}`,
                {
                    method: 'PUT',
                    body: JSON.stringify(settingsData),
                }
            );

            if (response.ok) {
                console.log('Data sent successfully');
                const updatedSettings: Settings = await response.json();
                setSettingsData(updatedSettings);
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <div className="w-full ">
            <h1 className="uppercase mb-5">Settings</h1>
            <form onSubmit={save} className="w-full max-w-lg">
                <div className="flex flex-wrap -mx-3 mb-6">
                    <div className="w-full px-3">
                        <label className="block uppercase tracking-wide text-white text-xs font-bold mb-2">
                            TELEGRAM_BOT_TOKEN
                        </label>
                        <input
                            name="telegramToken"
                            value={settingsData.telegramToken}
                            onChange={handleChange}
                            className="w-full appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                            type="text"
                            placeholder="<TELEGRAM_BOT_TOKEN>"
                        />
                        <p className="text-gray-600 text-xs italic underline hover:text-violet-600">
                            <a href="https://core.telegram.org/bots#how-do-i-create-a-bot">
                                How to get Bot Token?
                            </a>
                        </p>
                    </div>
                </div>

                <div className="flex flex-wrap -mx-3 mb-6">
                    <div className="w-full px-3">
                        <label className="block uppercase tracking-wide text-white text-xs font-bold mb-2">
                            TELEGRAM_CHAT_ID
                        </label>
                        <input
                            name="telegramChatId"
                            value={settingsData.telegramChatId}
                            onChange={handleChange}
                            className="w-full appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                            type="text"
                            placeholder="<TELEGRAM_CHAT_ID>"
                        />
                        <p className="text-gray-600 text-xs italic underline hover:text-violet-600">
                            <a href="https://stackoverflow.com/a/38388851">
                                How to get Chat ID?
                                https://stackoverflow.com/a/38388851
                            </a>
                        </p>
                        <p className="text-gray-600 text-xs italic hover:text-violet-600">
                            Example:
                            https://api.telegram.org/botTOKEN/getUpdates
                        </p>
                    </div>
                </div>

                {settingsData.id ? (
                    <button
                        type="submit"
                        onClick={update}
                        className="bg-violet-500 hover:bg-violet-600 text-white font-bold py-2 px-4 rounded flex w-200 items-center"
                    >
                        <div className="flex items-center">
                            <span>Update</span>
                        </div>
                    </button>
                ) : (
                    <button
                        type="submit"
                        onClick={save}
                        className="bg-violet-500 hover:bg-violet-600 text-white font-bold py-2 px-4 rounded flex w-200 items-center"
                    >
                        <div className="flex items-center">
                            <span>Save</span>
                        </div>
                    </button>
                )}
            </form>
        </div>
    );
}
