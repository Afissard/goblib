import type {Book} from "../types/book.ts";
import {type FormEvent, useEffect, useState} from "react";
import {updateBook} from "../services/api.ts";

interface Props {
    open: boolean;
    onClose: () => void;
    onUpdated?: () => void | Promise<void>;
    book?: Book;
}

const emptyForm = {
    title: "",
    author: "",
    summary: "",
    language: "",
    url: "",
    coverPath: ""
};

export default function BookUpdateForm(props: Props) {
    const { open, onClose, onUpdated, book } = props;
    const [form, setForm] = useState(emptyForm);
    const [isSaving, setIsSaving] = useState(false);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        if (open) {
            setForm({
                title: book?.title ?? "",
                author: book?.author ?? "",
                summary: book?.summary ?? "",
                language: book?.language ?? "",
                url: book?.url ?? "",
                coverPath: book?.coverPath ?? ""
            });
            setError(null);
            setIsSaving(false);
        }
    }, [open, book]);

    function close() {
        setForm(emptyForm);
        setError(null);
        setIsSaving(false);
        onClose();
    }

    if (!open) {
        return null;
    }

    async function handleSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();
        setError(null);
        setIsSaving(true);

        try {
            await updateBook({
                id: book?.id ?? "",
                title: form.title.trim(),
                author: form.author.trim() || "",
                summary: form.summary.trim() || "",
                language: form.language.trim() || "",
                url: form.url.trim() || "",
                coverPath: form.coverPath.trim()
            })

            setForm(emptyForm);
            await onUpdated?.();
            close();
        } catch (err) {
            setError(err instanceof Error ? err.message : "Failed to update book");
        } finally {
            setIsSaving(false);
        }
    }

    function updateField<K extends keyof typeof form>(field: K, value: string) {
        setForm((current) => ({
            ...current,
            [field]: value,
        }));
    }

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/70 p-4">
            <div
                className="w-full max-w-2xl rounded-lg border border-gray-700 bg-gray-900 p-6 shadow-2xl"
                onClick={(event) => event.stopPropagation()}
            >
                <div className="mb-6 flex items-center justify-between gap-4">
                    <div>
                        <h2 className="text-2xl font-bold">Update book</h2>
                        <p className="text-sm text-gray-400">Update an existing entry in the library.</p>
                    </div>

                    <button
                        type="button"
                        onClick={close}
                        className="rounded bg-gray-800 px-3 py-2 text-sm text-gray-200 hover:bg-gray-700"
                    >
                        Close
                    </button>
                </div>

                <form className="space-y-4" onSubmit={handleSubmit}>
                    <div>
                        <label className="mb-1 block text-sm font-medium text-gray-300" htmlFor="title">
                            Title
                        </label>
                        <input
                            id="title"
                            value={form.title}
                            onChange={(event) => updateField("title", event.target.value)}
                            className="w-full rounded border border-gray-700 bg-gray-800 px-3 py-2 text-white outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="The book title"
                            required
                        />
                    </div>

                    <div className="grid gap-4 md:grid-cols-2">
                        <div>
                            <label className="mb-1 block text-sm font-medium text-gray-300" htmlFor="author">
                                Author
                            </label>
                            <input
                                id="author"
                                value={form.author}
                                onChange={(event) => updateField("author", event.target.value)}
                                className="w-full rounded border border-gray-700 bg-gray-800 px-3 py-2 text-white outline-none focus:ring-2 focus:ring-blue-500"
                                placeholder="Author name"
                            />
                        </div>

                        <div>
                            <label className="mb-1 block text-sm font-medium text-gray-300" htmlFor="language">
                                Language
                            </label>
                            <input
                                id="language"
                                value={form.language}
                                onChange={(event) => updateField("language", event.target.value)}
                                className="w-full rounded border border-gray-700 bg-gray-800 px-3 py-2 text-white outline-none focus:ring-2 focus:ring-blue-500"
                                placeholder="English"
                            />
                        </div>
                    </div>

                    <div>
                        <label className="mb-1 block text-sm font-medium text-gray-300" htmlFor="url">
                            Source URL
                        </label>
                        <input
                            id="url"
                            value={form.url}
                            onChange={(event) => updateField("url", event.target.value)}
                            className="w-full rounded border border-gray-700 bg-gray-800 px-3 py-2 text-white outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="https://example.com/book"
                        />
                    </div>

                    <div>
                        <label className="mb-1 block text-sm font-medium text-gray-300" htmlFor="summary">
                            Summary
                        </label>
                        <textarea
                            id="summary"
                            value={form.summary}
                            onChange={(event) => updateField("summary", event.target.value)}
                            className="min-h-40 w-full rounded border border-gray-700 bg-gray-800 px-3 py-2 text-white outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="Short description of the book"
                        />
                    </div>

                    {error ? (
                        <div className="rounded border border-red-700 bg-red-950 px-3 py-2 text-sm text-red-300">
                            {error}
                        </div>
                    ) : null}

                    <div className="flex items-center justify-end gap-3 pt-2">
                        <button
                            type="button"
                            onClick={close}
                            className="rounded border border-gray-700 px-4 py-2 text-gray-200 hover:bg-gray-800"
                            disabled={isSaving}
                        >
                            Cancel
                        </button>
                        <button
                            type="submit"
                            className="rounded bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-500 disabled:cursor-not-allowed disabled:opacity-60"
                            disabled={isSaving}
                        >
                            {isSaving ? "Updating..." : "Update book"}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    )
}