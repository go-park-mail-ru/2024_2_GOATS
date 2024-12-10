CREATE TABLE public.subscriptions(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id int REFERENCES public.users(id),
  price int DEFAULT 0,
  status SUBSCRIPTION_TYPE_ENUM DEFAULT 'pending',
  expiration_date timestamp WITH TIME ZONE,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_subscriptions_expiration_date ON public.subscriptions(expiration_date);
CREATE INDEX idx_subscriptions_user_id ON public.subscriptions(user_id);
