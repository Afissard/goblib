
interface Props {
    onNew(): void;
}

export default function Toolbar({ onNew }: Props) {
    return (
        <div className="flex justify-between items-center p-4 border-b border-gray-700">
            <h1 className="text-3xl font-bold">
                Goblib
            </h1>

            <button
                onClick={onNew}
                className="bg-blue-600 hover:bg-blue-500 px-4 py-2 rounded"
            >
                New Book
            </button>
        </div>
    );
}