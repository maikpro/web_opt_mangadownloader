import { usePathname } from 'next/navigation';
import Link from 'next/link';
import Image from 'next/image';
import logo from '../../public/images/logo.png';

export default function Navbar() {
    const pathname = usePathname();
    return (
        <nav className="bg-white border-gray-200 dark:bg-gray-900">
            <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                <Link
                    href="/"
                    className="flex items-center space-x-3 rtl:space-x-reverse"
                >
                    <Image src={logo} alt="Logo" width={32} height={32} />
                    <h1>Web OPT Downloader</h1>
                </Link>

                <div className="w-full md:block md:w-auto" id="navbar-default">
                    <ul className="font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                        <li>
                            <Link
                                href="/"
                                className={`link ${
                                    pathname === '/'
                                        ? 'text-violet-600'
                                        : 'text-white'
                                }`}
                                aria-current="page"
                            >
                                Home
                            </Link>
                        </li>
                        <li>
                            <Link
                                href="/about"
                                className={`link ${
                                    pathname === '/about'
                                        ? 'text-violet-600'
                                        : 'text-white'
                                }`}
                            >
                                About
                            </Link>
                        </li>
                        <li>
                            <Link
                                href="/settings"
                                className={`link ${
                                    pathname === '/settings'
                                        ? 'text-violet-600'
                                        : 'text-white'
                                }`}
                            >
                                Settings
                            </Link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    );
}
