-- Заполнение таблицы questions
INSERT INTO public.questions (question, is_active) VALUES
                                                       ('Как вы оцениваете наш сервис?', TRUE),
                                                       ('Насколько вы довольны качеством обслуживания?', TRUE),
                                                       ('Порекомендуете ли вы нас друзьям?', FALSE);

-- Заполнение таблицы answers
INSERT INTO public.answers (question_id, answer) VALUES
                                                     (1, 'Отлично'),
                                                     (1, 'Хорошо'),
                                                     (1, 'Удовлетворительно'),
                                                     (1, 'Плохо'),
                                                     (2, 'Очень доволен'),
                                                     (2, 'Доволен'),
                                                     (2, 'Не доволен'),
                                                     (3, 'Да'),
                                                     (3, 'Нет');

-- Заполнение таблицы user_feedback_status
INSERT INTO public.user_feedback_status (user_id, passed) VALUES
                                                              (1, TRUE),
                                                              (2, FALSE),
                                                              (3, TRUE);

-- Заполнение таблицы survey_data
INSERT INTO public.survey_data (user_id, question_id, answer_id, answer) VALUES
                                                                             (1, 1, 1, NULL),
                                                                             (1, 2, 5, NULL),
                                                                             (2, 1, 3, NULL),
                                                                             (3, 1, NULL, 'Ваш сервис лучший!'),
                                                                             (3, 2, NULL, 'Доволен на 90%');