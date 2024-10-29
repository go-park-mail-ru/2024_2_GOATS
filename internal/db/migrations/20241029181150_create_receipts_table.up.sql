CREATE TABLE public.receipts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    payment_id BIGINT NOT NULL REFERENCES public.payments(id),
    receipt_type RECEIPT_TYPE_ENUM NOT NULL DEFAULT 'created',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_receipts_payment_id ON public.receipts(payment_id);
CREATE INDEX idx_receipts_receipt_type ON public.receipts(receipt_type);
