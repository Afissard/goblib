interface Props {
    book?: any;
}

export default function BookDetails({ book }: Props) {
    if (!book) {
        return (
            <div className="h-full flex items-center justify-center text-gray-400">
                Select a book to see details
            </div>
        );
    }

    const { title, author, summary, language, source_url } = book // TODO add cover

    return (
        <div className="max-w-4xl mx-auto bg-gray-800 rounded p-6">
            <div className="flex gap-6">
                <img src="../assets/hero.png" alt="cover" className="w-48 h-64 object-cover rounded" />
                <div className="flex-1">
                    <div className="text-2xl font-bold">{title}</div>
                    <div className="text-sm text-gray-400 mb-4">{author}</div>
                    <div className="text-sm text-gray-300 mb-4"><strong>Language:</strong> {language}</div>
                    <div className="text-sm text-gray-200 mb-4">{summary}</div>
                    {source_url ? (
                        <a className="text-sm text-blue-400 hover:underline" href={source_url} target="_blank" rel="noreferrer">
                            Open source
                        </a>
                    ) : null}
                </div>
            </div>
        </div>
    );
}