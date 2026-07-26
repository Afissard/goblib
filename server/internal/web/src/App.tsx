import { useEffect, useState } from "react";
import BookList from "./components/BookList";
import BookDetails from "./components/BookDetails";
import { listBooks } from "./services/api";
import Toolbar from "./components/Toolbar.tsx";
import BookForm from "./components/BookForm";
import BookUpdateForm from "./components/BookUpdateForm.tsx";

function normalize(book: any) {
    return {
        id: book.id ?? book.ID ?? "",
        title: book.title ?? book.Title ?? "Untitled",
        author: book.author ?? book.Author ?? "",
        summary: book.summary ?? book.Summary ?? "",
        language: book.language ?? book.Language ?? "",
        source_url: book.source_url ?? book.SourceURL ?? "",
        cover_path: book.cover_path ?? book.CoverPath ?? "",
        raw: book,
    };
}

export default function App() {
    const [books, setBooks] = useState<any[]>([]);
    const [selected, setSelected] = useState<any | undefined>(undefined);
    const [isBookFormOpen, setIsBookFormOpen] = useState(false);
    const [isBookUpdateFormOpen, setIsBookUpdateFormOpen] = useState(false);

    async function refreshBooks() {
        try {
            const data = await listBooks();
            setBooks(data);
            console.log("Books loaded:", data);
        } catch (err) {
            console.error("Failed to load books:", err);
        }
    }

    useEffect(() => {
        refreshBooks();
    }, []);

    function handleSelect(book: any) {
        const n = normalize(book);
        console.log("Selected:", n);
        setSelected(n);
    }

    return (
        <div className="h-screen flex flex-col bg-gray-900 text-white">
            <Toolbar onNew={() => setIsBookFormOpen(true)} />
            <div className="flex flex-1">
                <div className="w-72 border-r border-gray-700">
                    <BookList books={books} selected={selected} onSelect={handleSelect} />
                </div>

                <div className="flex-1 p-6 overflow-auto">
                    <BookDetails book={selected} onUpdateBook={ () => setIsBookUpdateFormOpen(true) } />
                </div>
            </div>

            <BookForm
                open={isBookFormOpen}
                onClose={() => setIsBookFormOpen(false)}
                onCreated={refreshBooks}
            />

            <BookUpdateForm
                open={isBookUpdateFormOpen}
                onClose={() => setIsBookUpdateFormOpen(false)}
                onUpdated={refreshBooks}
                book={selected}
            />
        </div>
    );
}