import { useEffect, useState } from "react";

import Toolbar from "./components/Toolbar";
import BookDetails from "./components/BookDetails";
import BookForm from "./components/BookForm";
import BookList from "./components/BookList";

import {
    listBooks,
    getBook,
    createBook,
    updateBook,
    deleteBook,
} from "./services/api";

import type {Book} from "./types/book";

export default function App() {

    const [books, setBooks] = useState<Book[]>([]);
    const [selected, setSelected] = useState<Book>();
    const [editing, setEditing] = useState(false);

    async function refresh() {
        const data = await listBooks();
        setBooks(data);
    }

    useEffect(() => {
        refresh();
    }, []);

    async function selectBook(book: Book) {
        const full = await getBook(book.id);
        setSelected(full);
        setEditing(false);
    }

    async function saveBook(book: Partial<Book>) {

        if (editing && selected) {

            await updateBook({
                ...selected,
                ...book,
            });

        } else {

            await createBook(book);

        }

        await refresh();

        setEditing(false);

        setSelected(undefined);
    }

    async function removeBook() {

        if (!selected)
            return;

        await deleteBook(selected.id);

        await refresh();

        setSelected(undefined);
    }

    return (
        <div className="h-screen flex flex-col bg-gray-900 text-white">

            <Toolbar
                onNew={() => {
                    setSelected(undefined);
                    setEditing(true);
                }}
            />

            <div className="flex flex-1">

                <div className="w-72 border-r border-gray-700">

                    <BookList
                        books={books}
                        selected={selected}
                        onSelect={selectBook}
                    />

                </div>

                <div className="flex-1">

                    {editing ? (

                        <BookForm
                            initial={selected}
                            onSubmit={saveBook}
                        />

                    ) : (

                        <BookDetails
                            book={selected}
                            onEdit={() => setEditing(true)}
                            onDelete={removeBook}
                        />

                    )}

                </div>

            </div>

        </div>
    );
}