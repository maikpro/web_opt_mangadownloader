'use client';
import Checkbox from '@/components/checkbox/checkbox';
import { Arc } from '@/models/Arc';
import { ArcService } from '@/services/arc-service';
import { ChangeEvent, SetStateAction, useEffect, useState } from 'react';

export default function Home() {
    const [arcs, setArcs] = useState<Arc[]>([]);
    const [selectedChapter, setSelectedChapter] = useState('');
    const [isDownload, setIsDownload] = useState<boolean>(false);
    const [savedPath, setSavedPath] = useState<any>('');
    const [isTelegram, setIsTelegram] = useState<boolean>(true);
    const [isLocal, setLocal] = useState<boolean>(true);

    useEffect(() => {
        const getArcs = async () => {
            const arcs = await ArcService.getArcs();
            setArcs(arcs);
            setSelectedChapter('' + arcs[0].entries[0].number);
        };
        getArcs();
    }, []);

    const handleTelegramChange = (checked: boolean) => {
        setIsTelegram(checked);
    };

    const handleLocalChange = (checked: boolean) => {
        setLocal(checked);
    };

    const handleSelectChange = (event: {
        target: { value: SetStateAction<string> };
    }) => {
        setSelectedChapter(event.target.value);
    };

    const handleDownload = async () => {
        setIsDownload(true);
        if (!selectedChapter) {
            console.error('No Chapter selected...');
            return;
        }
        console.log(
            `You have selected chapter '${selectedChapter}', isLocal '${isLocal}', isTelegram ${isTelegram} `
        );
        const response = await fetch(
            `${process.env.NEXT_PUBLIC_API_URL}/chapters/id/${selectedChapter}?local=${isLocal}&telegram=${isTelegram}`,
            { method: 'POST' }
        );

        if (response.ok) {
            setIsDownload(false);
        }

        if (response.ok && isLocal) {
            const path = await response.json();
            console.log(path);
            setSavedPath(path);
            if (path) {
                setIsDownload(false);
            }
        }
    };

    return (
        <div>
            {arcs.length > 0 ? (
                <div className="flex flex-col space-y-4 flex-wrap">
                    <h1>Select a chapter to download:</h1>

                    <select
                        className="row-start-1 col-start-1 bg-slate-50 dark:bg-slate-800 p-3 rounded"
                        value={selectedChapter}
                        onChange={handleSelectChange}
                    >
                        {arcs.map((arc, index) => (
                            <optgroup label={arc.name} key={index}>
                                {arc.entries.map((entry) => (
                                    <option
                                        key={entry.id}
                                        value={entry.number}
                                        disabled={!entry.is_available}
                                    >
                                        {entry.number} - {entry.name}
                                    </option>
                                ))}
                            </optgroup>
                        ))}
                    </select>
                    <div className="w-full flex justify-between">
                        <Checkbox
                            labelText="Download on your local drive?"
                            isChecked={isLocal}
                            onChange={handleLocalChange}
                            isDisabled={isDownload}
                        />

                        <Checkbox
                            labelText="Send to your Checked Chat?"
                            isChecked={isTelegram}
                            onChange={handleTelegramChange}
                            isDisabled={isDownload}
                        />
                    </div>
                    <button
                        onClick={handleDownload}
                        className={`w-full max-w-40 bg-violet-500 hover:bg-violet-600 text-white font-bold py-2 px-4 rounded flex w-200 justify-center items-center ${
                            isDownload ? 'opacity-50 cursor-not-allowed' : ''
                        }`}
                        disabled={isDownload}
                    >
                        {isDownload ? (
                            <div className="flex items-center">
                                <svg
                                    className="animate-spin border-white-400 rounded-full h-4 w-4 border-b-2 mr-2"
                                    viewBox="0 0 20 20"
                                ></svg>
                                <span>Processing...</span>
                            </div>
                        ) : (
                            <div className="flex items-center">
                                <svg
                                    className="fill-current w-4 h-4 mr-2"
                                    xmlns="http://www.w3.org/2000/svg"
                                    viewBox="0 0 20 20"
                                >
                                    <path d="M13 8V2H7v6H2l8 8 8-8h-5zM0 18h20v2H0v-2z" />
                                </svg>
                                <span>Download</span>
                            </div>
                        )}
                    </button>
                </div>
            ) : (
                <div className="flex items-center">
                    <p className="text-red-500">
                        <svg
                            className="animate-spin border-white-400 rounded-full h-4 w-4 border-b-2 mr-2"
                            viewBox="0 0 20 20"
                        ></svg>
                        fetching chapters from onepiece-tube.com...
                    </p>
                </div>
            )}

            <p className="text-green-500">{savedPath.path}</p>
        </div>
    );
}
