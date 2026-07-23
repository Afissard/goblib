import type {Book} from "../types/book";

const API = "/api/books";

export interface CreateBookInput {
    title: string;
    author?: string;
    summary?: string;
    language?: string;
    source_url?: string;
}

async function request<T>(url: string, init?: RequestInit): Promise<T> {
    const response = await fetch(url, init);

    if (!response.ok) {
        throw new Error(await response.text());
    }

    return response.json();
}

export function listBooks(): Promise<Book[]> {
    return request<Book[]>(API);
}

export function getBook(id: string): Promise<Book> {
    return request<Book>(`${API}/${id}`);
}

export async function createBook(book: CreateBookInput): Promise<void> {
    const params = new URLSearchParams();

    params.set("title", book.title);

    if (book.author) params.set("author", book.author);
    if (book.summary) params.set("summary", book.summary);
    if (book.language) params.set("language", book.language);
    if (book.source_url) params.set("source_url", book.source_url);

    const response = await fetch(`${API}?${params.toString()}`, {
        method: "POST",
    });

    if (!response.ok) {
        throw new Error(await response.text());
    }
}

export function updateBook(book: Book): Promise<Book> {
    return request<Book>(`${API}/${book.id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(book),
    });
}

export async function deleteBook(id: string): Promise<void> {
    const response = await fetch(`${API}/${id}`, {
        method: "DELETE",
    });

    if (!response.ok) {
        throw new Error(await response.text());
    }
}