export default function Settings() {
    const save = () => {
        console.log('Save changes...');
    };

    return (
        <div className="w-full ">
            <h1 className="uppercase mb-5">Settings</h1>
            <form className="w-full max-w-lg">
                <div className="flex flex-wrap -mx-3 mb-6">
                    <div className="w-full px-3">
                        <label className="block uppercase tracking-wide text-white text-xs font-bold mb-2">
                            TELEGRAM_BOT_TOKEN
                        </label>
                        <input
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

                <button
                    onClick={save}
                    className="bg-violet-500 hover:bg-violet-600 text-white font-bold py-2 px-4 rounded flex w-200 items-center"
                >
                    <div className="flex items-center">
                        <span>Save</span>
                    </div>
                </button>
            </form>
        </div>
    );
}
