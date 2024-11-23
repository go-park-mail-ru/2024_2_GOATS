-- Таблица для хранения данных опросов
CREATE TABLE public.survey_data (
        id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, -- Уникальный идентификатор записи
        user_id BIGINT NOT NULL, -- Идентификатор пользователя
        question_id BIGINT NOT NULL, -- Идентификатор вопроса
        answer_id BIGINT, -- Идентификатор выбранного ответа (для радиобатонов)
        answer TEXT, -- Текст ответа (для текстовых вопросов)
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Дата создания записи
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_survey_data_user_id ON public.survey_data(user_id);
CREATE INDEX idx_survey_data_question_id ON public.survey_data(question_id);

-- Таблица для хранения вопросов
CREATE TABLE public.questions (
      question_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, -- Уникальный идентификатор вопроса
      question TEXT NOT NULL, -- Текст вопроса
      is_active BOOLEAN NOT NULL DEFAULT TRUE, -- Флаг активности вопроса
      created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_questions_is_active ON public.questions(is_active);

-- Таблица для хранения вариантов ответа
CREATE TABLE public.answers (
    answer_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, -- Уникальный идентификатор ответа
    question_id BIGINT NOT NULL REFERENCES public.questions(question_id) ON DELETE CASCADE, -- Связь с вопросом
    answer TEXT NOT NULL, -- Текст ответа
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_answers_question_id ON public.answers(question_id);

-- Таблица для отслеживания прохождения опроса пользователем
CREATE TABLE public.user_feedback_status (
     id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, -- Уникальный идентификатор записи
     user_id BIGINT NOT NULL, -- Идентификатор пользователя
     passed BOOLEAN NOT NULL DEFAULT FALSE, -- Флаг прохождения опроса
     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

     UNIQUE (user_id)
);

CREATE INDEX idx_user_feedback_status_user_id ON public.user_feedback_status(user_id);
