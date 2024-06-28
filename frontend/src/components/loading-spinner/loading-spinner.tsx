interface LoadingSpinnerProps {
    loadingText: string;
}

export default function LoadingSpinner(
    loadingSpinnerProps: LoadingSpinnerProps
) {
    return (
        <div className="flex items-center">
            <p className="text-red-500">
                <svg
                    className="animate-spin border-white-400 rounded-full h-4 w-4 border-b-2 mr-2"
                    viewBox="0 0 20 20"
                ></svg>
                {loadingSpinnerProps.loadingText}
            </p>
        </div>
    );
}
