'use client';
import { SetStateAction, useEffect, useState } from 'react';

import { Arc } from './models/Arc';

export default function Home() {
    const [arcs, setArcs] = useState<Arc[]>([]);
    const [selectedChapter, setSelectedChapter] = useState('');
    const [isDownload, setIsDownload] = useState<boolean>(false);
    const [savedPath, setSavedPath] = useState<any>('');

    useEffect(() => {
        const fetchArcs = async () => {
            const response = await fetch('http://localhost:8080/api/arcs');
            const arcs: Arc[] = await response.json();
            setArcs(arcs);
            setSelectedChapter('' + arcs[0].entries[0].number);
        };
        fetchArcs();
    }, []);

    // Handler function to update the selected value
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
        console.log('You have selected: ', selectedChapter);
        const response = await fetch(
            `http://localhost:8080/api/chapters/id/${selectedChapter}`,
            { method: 'POST' }
        );
        const path = await response.json();
        console.log(path);
        setSavedPath(path);
        if (path) {
            setIsDownload(false);
        }
    };

    return (
        <main className="flex flex-col space-y-4 flex-wrap">
            <h1>Web OPT Downloader</h1>

            <select
                className="appearance-none row-start-1 col-start-1 bg-slate-50 dark:bg-slate-800"
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

            <p>Selected value: {selectedChapter}</p>

            <button
                onClick={handleDownload}
                className={`w-full max-w-40 bg-gray-400 hover:bg-gray-500 text-gray-800 font-bold py-2 px-4 rounded flex w-200 items-center ${
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

            <p className="text-green-500">{savedPath.path}</p>
        </main>
    );
}
