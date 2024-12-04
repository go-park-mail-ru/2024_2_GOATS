DROP TABLE IF EXISTS payments CASCADE;

CREATE TABLE public.payments(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  subscription_id int,
  requested_amount int DEFAULT 0,
  captured_amount int DEFAULT 0,
  captured_at timestamp WITH TIME ZONE,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_payments_subscription_id ON public.payments(subscription_id);
CREATE INDEX idx_payments_captured_at ON public.payments(captured_at);
CREATE INDEX idx_payments_created_at ON public.payments(created_at);
CREATE INDEX idx_payments_updated_at ON public.payments(updated_at);
