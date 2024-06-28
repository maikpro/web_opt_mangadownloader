'use client';

import ChapterPicker from '@/components/chapter-picker/chapter-picker';
import Reddit from '@/components/reddit/reddit';

export default function Home() {
    return (
        <div>
            <Reddit />
            <ChapterPicker />
        </div>
    );
}
