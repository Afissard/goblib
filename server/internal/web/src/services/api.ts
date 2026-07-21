import type {Book} from "../types/book";

const API = "/api/books";

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

export function createBook(book: Partial<Book>): Promise<Book> {
    return request<Book>(API, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(book),
    });
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