CREATE TABLE public.payments (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    subscription_id BIGINT NOT NULL REFERENCES public.subscriptions(id),
    captured_total DECIMAL(10, 2) NOT NULL DEFAULT '0.0',
    refunded_total DECIMAL(10, 2) NOT NULL DEFAULT '0.0',
    payment_number INT NOT NULL DEFAULT 1,
    status PAYMENT_STATUS_ENUM NOT NULL DEFAULT 'started',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_payments_subscription_id ON public.payments(subscription_id);
CREATE INDEX idx_payments_status ON public.payments(status);
CREATE INDEX idx_payments_payment_number ON public.payments(payment_number);
