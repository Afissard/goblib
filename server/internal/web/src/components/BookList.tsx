import BookLine from "./BookLine";

interface Props {
    books?: any[];
    selected?: any;
    onSelect?(book: any): void;
}

export default function BookList({ books = [], selected, onSelect }: Props) {
    if (!books || books.length === 0) {
        return <div className="p-4 text-gray-400">No books found</div>;
    }
    // TODO: add sorting by name, author, etc
    return (
        <div className="overflow-auto">
            {books.map((book, idx) => (
                <BookLine
                    key={book.id ?? book.ID ?? idx}
                    book={book}
                    selected={selected}
                    onSelect={onSelect}
                />
            ))}
        </div>
    );
}