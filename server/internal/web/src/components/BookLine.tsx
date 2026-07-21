interface Props {
    book: any;
    selected?: any;
    onSelect?(book: any): void;
}

export default function BookLine({ book, selected, onSelect }: Props) {
    const id = book?.id ?? book?.ID ?? "";
    const title = book?.title ?? book?.Title ?? "Untitled";
    const author = book?.author ?? book?.Author ?? "";

    const selectedId = selected ? (selected.id ?? selected.ID ?? undefined) : undefined;
    return (
        <div
            onClick={() => onSelect?.(book)}
            className={
                "cursor-pointer p-3 border-b border-gray-700 hover:bg-gray-800 " +
                (selectedId === id ? "bg-gray-800" : "")
            }
        >
            <div className="font-semibold">{title}</div>
            <div className="text-sm text-gray-400">{author}</div>
        </div>
    );
}