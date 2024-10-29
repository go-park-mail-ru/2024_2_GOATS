CREATE TABLE public.line_items (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    receipt_id BIGINT NOT NULL REFERENCES public.receipts(id),
    title TEXT NOT NULL,
    total DECIMAL(10, 2) NOT NULL DEFAULT '0.0',
    line_item_type LINE_ITEM_TYPE_ENUM NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_line_items_receipt_id ON public.line_items(receipt_id);
CREATE INDEX idx_line_items_line_item_type ON public.line_items(line_item_type);
