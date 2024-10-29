CREATE TABLE public.subscriptions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES public.users(id),
    price DECIMAL(10, 2) NOT NULL DEFAULT '0.0',
    start_date DATE NOT NULL,
    days_counter INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_subscriptions_user_id ON public.subscriptions(user_id);
CREATE INDEX idx_subscriptions_start_date ON public.subscriptions(start_date);
