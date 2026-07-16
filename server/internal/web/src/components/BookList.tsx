import type {Book} from "../types/book";

interface Props {
    books: Book[];
    selected?: Book;
    onSelect(book: Book): void;
}

export default function BookList({
                                     books,
                                     selected,
                                     onSelect
                                 }: Props) {

    return (
        <div className="overflow-auto">

            {books.map(book => (

                <div
                    key={book.id}
                    onClick={() => onSelect(book)}
                    className={
                        "cursor-pointer p-3 border-b border-gray-700 hover:bg-gray-800 " +
                        (selected?.id === book.id ? "bg-gray-800" : "")
                    }
                >
                    <div className="font-semibold">
                        {book.title}
                    </div>

                    <div className="text-sm text-gray-400">
                        {book.author}
                    </div>

                </div>

            ))}

        </div>
    );
}