import { useState } from "react";
import type {Book} from "../types/book";

interface Props {
    initial?: Partial<Book>;
    onSubmit(book: Partial<Book>): void;
}

export default function BookForm({ initial, onSubmit }: Props) {

    const [title, setTitle] = useState(initial?.title ?? "");
    const [author, setAuthor] = useState(initial?.author ?? "");

    return (

        <div className="space-y-4">

            <input
                className="w-full bg-gray-800 p-2 rounded"
                placeholder="Title"
                value={title}
                onChange={e => setTitle(e.target.value)}
            />

            <input
                className="w-full bg-gray-800 p-2 rounded"
                placeholder="Author"
                value={author}
                onChange={e => setAuthor(e.target.value)}
            />

            <button
                className="bg-green-700 px-4 py-2 rounded"
                onClick={() =>
                    onSubmit({
                        ...initial,
                        title,
                        author
                    })
                }
            >
                Save
            </button>

        </div>

    );
}