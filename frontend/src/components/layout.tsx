import Navbar from './nav/navbar';

export default function Layout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <>
            <Navbar />
            <main className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                {children}
            </main>
        </>
    );
}
