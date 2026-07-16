import type {Book} from "../types/book";

interface Props {

    book?: Book;

    onEdit(): void;
    onDelete(): void;
}

export default function BookDetails({
                                        book,
                                        onEdit,
                                        onDelete
                                    }: Props) {

    if (!book)
        return (
            <div className="p-6">
                Select a book.
            </div>
        );

    return (

        <div className="p-6 space-y-4">

            <h2 className="text-2xl font-bold">
                {book.title}
            </h2>

            <p>
                <strong>Author:</strong> {book.author}
            </p>

            <p>
                <strong>Language:</strong> {book.language}
            </p>

            <p>
                {book.summary}
            </p>

            <div className="flex gap-3">

                <button
                    className="bg-yellow-600 px-4 py-2 rounded"
                    onClick={onEdit}
                >
                    Edit
                </button>

                <button
                    className="bg-red-700 px-4 py-2 rounded"
                    onClick={onDelete}
                >
                    Delete
                </button>

            </div>

        </div>

    );
}