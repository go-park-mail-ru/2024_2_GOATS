INSERT INTO public.genres (title)
VALUES
  ('Ужасы'),
  ('Комедия'),
  ('Психология'),
  ('Фэнтези'),
  ('Триллер'),
  ('Романтика'),
  ('Спорт'),
  ('Драма');

INSERT INTO public.countries (title, code)
VALUES
  ('Россия', 'rus'),
  ('США', 'usa'),
  ('Великобритания', 'uk'),
  ('Южная Корея', 'kor'),
  ('Испания', 'es'),
  ('Франция', 'fr'),
  ('Дания', 'dk');

INSERT INTO public.movies (video_url, title_url, title, short_description, card_url, album_url, country_id, rating, long_description, release_date, movie_type)
VALUES
  ('/static/movies/squid-game/movie.mp4', '/static/movies/squid-game/logo.png', 'Игра в кальмара', 'Под маской детских игр и забавных заданий скрывается жестокая конкуренция, которая приводит к драматическим последствиям.', '/static/movies/squid-game/card.png', '/static/movies/squid-game/poster.png', 4, 7.6, 'Сон Ги-хун уже немолод, разведён, по уши погряз в долгах и сидит на шее у старенькой матери. Даже выигранные на скачках деньги в его руках долго не задерживаются, и однажды он встречает в метро загадочного незнакомца, который сначала предлагает сыграть в детскую игру, а затем вручает Ги-хуну немалую сумму и визитку. Но радость мужчины сменятся отчаянием, когда он узнаёт, что бывшая жена с новым мужем собираются увезти его дочь в Америку. Он звонит по номеру с визитки и становится последним участником тайных игр на выживание с призом в 40 миллионов долларов. Среди товарищей по несчастью оказываются его друг детства — прогоревший финансист, бандит, смертельно больной старик, северокорейская перебежчица, иммигрант из Пакистана и многие другие отчаянно нуждающиеся в деньгах.', '2021-09-17', 'serial'),
  ('/static/movies/la-casa-de-papel/movie.mp4', '/static/movies/la-casa-de-papel/logo.png', 'Бумажный дом', 'короткое описание', '/static/movies/la-casa-de-papel/card.png', '/static/movies/la-casa-de-papel/poster.png', 5, 8.2, 'История о преступниках, решивших ограбить Королевский монетный двор Испании и украсть 2,4 млрд евро.', '2017-05-02', 'serial'),
  ('/static/movies/better-call-saul/movie.mp4', '/static/movies/better-call-saul/logo.png', 'Лучше звоните Соулу', 'короткое описание', '/static/movies/better-call-saul/card.png', '/static/movies/better-call-saul/poster.png', 2, 8.4, 'История об испытаниях и невзгодах, которые приходится преодолеть Солу Гудману, адвокату по уголовным делам, в тот период, когда он пытается открыть свою собственную адвокатскую контору в Альбукерке, штат Нью-Мексико.', '2015-02-08', 'serial'),

  ('/static/movies/1+1/movie.mp4', '/static/movies/1+1/logo.png', '1+1', 'Аристократ на коляске нанимает в сиделки бывшего заключенного. Искрометная французская комедия с Омаром Си', '/static/movies/1+1/card.png', '/static/movies/1+1/poster.png', 6, 8.8, 'Жизнеутверждающая комедия о дружбе и преодолении трудностей. Филиппу требуется уход, и из всех претендентов на должность помощника он выбирает молодого сенегальца Дрисса с криминальным прошлым. Несмотря на расовые и классовые различия, Филипп и Дрисс становятся настоящими друзьями и обретают новый опыт. Удивительно воодушевляющая история, основанная на реальных событиях.', '2011-09-23', 'film'),
  ('/static/movies/avatar/movie.mp4', '/static/movies/avatar/logo.png', 'Аватар', 'Парализованный морпех становится мессией для жителей Пандоры. Самый кассовый фильм в истории кино', '/static/movies/avatar/card.png', '/static/movies/avatar/poster.png', 2, 8.0, 'Бывший морпех Джейк Салли прикован к инвалидному креслу. Несмотря на немощное тело, Джейк в душе по-прежнему остается воином. Он получает задание совершить путешествие в несколько световых лет к базе землян на планете Пандора, где корпорации добывают редкий минерал, имеющий огромное значение для выхода Земли из энергетического кризиса.', '2009-12-10', 'film'),
  ('/static/movies/how-you-see-me/movie.mp4', '/static/movies/how-you-see-me/logo.png', 'Иллюзия обмана', 'Банда трюкачей разоблачает фокусы коррупционеров. Ошеломляющие магические ограбления от Джесси Айзенберга и компании', '/static/movies/how-you-see-me/card.png', '/static/movies/how-you-see-me/poster.png', 2, 7.7, 'Зрелищный экшн-триллер о талантливых иллюзионистах, играющих не только с законами разума. Что может быть более волнующим чем ограбление банка прямо на сцене телешоу, особенно когда полученные деньги сыпятся прямо в зрительный зал. Именно такими номерами славится великолепная четвёрка иллюзионистов, объединенная таинственным незнакомцем. Удастся ли заинтересовавшимся шоу властям разгадать секрет номера? В картине приняли участия такие звёзды кино как: Марк Руффало, Морган Фриман, Вуди Харрельсон, Майкл Кейн и другие.', '2013-05-21', 'film'),
  ('/static/movies/inglorious-basterds/movie.mp4', '/static/movies/inglorious-basterds/logo.png', 'Бесславные ублюдки', 'Американский спецотряд жестоко расправляется с нацистами. Пародия на военные фильмы от Квентина Тарантино', '/static/movies/inglorious-basterds/card.png', '/static/movies/inglorious-basterds/poster.png', 2, 8.0, 'Провокационная комедия Квентина Тарантино о специальном отряде солдат, охотящихся на нацистов во время Второй мировой войны. Поклонники фирменного стиля режиссера будут рады: в фильме есть и юмор, и пародия на киножанры, и интеллектуальные задачки для любителей кино. Актерский состав на высоте. Например, игра Кристофа Вальца в этом фильме принесла ему первый «Оскар» за мужскую роль.', '2009-05-20', 'film'),
  ('/static/movies/interstellar/movie.mp4', '/static/movies/interstellar/logo.png', 'Интерстеллар', 'Фантастический эпос про задыхающуюся Землю, космические полеты и парадоксы времени. «Оскар» за спецэффекты', '/static/movies/interstellar/card.png', '/static/movies/interstellar/poster.png', 2, 8.6, 'Когда засуха приводит человечество к продовольственному кризису, коллектив исследователей и учёных отправляется сквозь червоточину (которая предположительно соединяет области пространства-времени через большое расстояние) в путешествие, чтобы превзойти прежние ограничения для космических путешествий человека и переселить человечество на другую планету.', '2014-11-06', 'film'),
  ('/static/movies/legend17/movie.mp4', '/static/movies/legend17/logo.png', 'Легенда №17', 'В хоккей играют настоящие герои. Российский байопик о звезде советского хоккея Валерии Харламове с Данилой Козловским', '/static/movies/legend17/card.png', '/static/movies/legend17/poster.png', 1, 8.0, 'Российский байопик о звезде советского хоккея. 1972 год стал для советского хоккея запоминающимся. Знаменательный матч между командами СССР и Канады вспоминают до сих пор. «Легенда № 17» (2013) — это не просто кино, это дань уважения величайшим спортсменам прошлого. Валерий Харламов влюбился в хоккей еще в раннем возрасте и постепенно достигает больших успехов. Вот только травма способна разрушить все планы на будущее. Отчаявшись, парень опускает руки, «ломается». Только благодаря упорству, профессионализму и жесткости тренера, Анатолия Тарасова, он вновь выходит на лед. Именно его возвращение в русскую сборною сделало результат матча. Хотите узнать, каково это было? Смотрите фильм онлайн в нашем кинотеатре.', '2013-04-18', 'film'),
  ('/static/movies/moneyball/movie.mp4', '/static/movies/moneyball/logo.png', 'Человек, который изменил все', 'Менеджер-визионер пытается сделать из аутсайдеров чемпионов. Спортивная драма по сценарию Аарона Соркина', '/static/movies/moneyball/card.png', '/static/movies/moneyball/poster.png', 2, 7.7, '2001 год. Менеджер бейсбольной команды Oakland Athletics Билли Бин после проигрыша в ключевом матче команде New York Yankees, которая обладает значительно большим бюджетом, решает в корне изменить систему подбора игроков. Во время деловой поездки в Кливленд он знакомится с молодым выпускником Йеля, экономистом Питером Брэндом, который при помощи математических расчётов предлагает внедрить новаторскую схему оценки полезности игроков, основываясь на показателях их личной статистики.', '2011-09-09', 'film'),
  ('/static/movies/shutter-island/movie.mp4', '/static/movies/shutter-island/logo.png', 'Остров проклятых', 'Тонкая грань между реальностью и безумием в детективе от Мартина Скорсезе и разобранный на мемы ДиКаприо', '/static/movies/shutter-island/card.png', '/static/movies/shutter-island/poster.png', 2, 8.5, '1950-е. Федеральный маршал Эдвард Дэниелс (Леонардо ДиКаприо) отправляется на неприступный остров в особо охраняемую лечебницу для душевнобольных преступников. Ему с напарником предстоит расследовать исчезновение из закрытой камеры убийцы троих детей Рейчел Соландо. Мрачный остров, таящий множество тайн, откроет перед героем страшные подробности его прошлого.', '2010-02-13', 'film'),
  ('/static/movies/taxi2/movie.mp4', '/static/movies/taxi2/logo.png', 'Такси 2', 'Устроить роды в машине, успеть на званый обед и разобраться с якудзой — очередные приключения водителя белого «Пежо»', '/static/movies/taxi2/card.png', '/static/movies/taxi2/poster.png', 6, 7.7, 'Пристегнуть ремни безопасности хотят даже зрители в кинотеатре. Неугомонный Даниэль по-прежнему работает таксистом в Марселе и, несмотря на бесчисленное количество штрафов, придерживается совершенно бешеного стиля езды. Только теперь он еще старается построить отношения с любимой девушкой Лили, что значит, понравиться и ее родителям. Неожиданно суровому отцу Лили, генералу Бертенье, понадобился водитель с машиной. Будущий зятек смог показать себя во всей красе на скорости 300 км\час. Так, благодаря рекомендациям генерала, простой таксист становится водителем спецмашины министра Франции.', '2000-03-25', 'film'),
  ('/static/movies/the-many-saints-of-newark/movie.mp4', '/static/movies/the-many-saints-of-newark/logo.png', 'Множественные святые Ньюагка', 'Итальянская мафия борется за власть в Америке 1960-70-х. Приквел «Сопрано» по сценарию Дэвида Чейза', '/static/movies/the-many-saints-of-newark/card.png', '/static/movies/the-many-saints-of-newark/poster.png', 2, 5.9, '1967 год, Ньюарк в штате Нью-Джерси. Маленький Тони Сопрано восхищается деловым партнёром отца — Ричардом Молтисанти по прозвищу Дики. Тот ответственно ведёт доверенную ему отцом, местным криминальным авторитетом, часть бизнеса и, в отличие от остальных мафиози, не гнушается иметь дело с чернокожими. Когда в городе вспыхивают беспорядки, спровоцированные жестоким обращением полиции с чёрным таксистом, Дики совершает импульсивный поступок и использует окружающий хаос в качестве прикрытия.', '2021-09-22', 'film'),
  ('/static/movies/the-transporter/movie.mp4', '/static/movies/the-transporter/logo.png', 'Перевозчик', 'Перевозя очередную посылку, Фрэнк вынужден нарушить привычные правила. Джейсон Стэйтем в скоростном роуд-муви', '/static/movies/the-transporter/card.png', '/static/movies/the-transporter/poster.png', 2, 7.4, 'Бывший десантник Фрэнк Мартин имеет неплохой бизнес - перевозит любые грузы по французскому Средиземноморью и делает свою работу быстро и качественно. Недостатка в клиентах нет, ведь он всегда неукоснительно соблюдает три правила: не меняет условий сделки, не спрашивает никаких имен и никогда не заглядывает в багаж. Но однажды, перевозя груз клиента по имени Уолл Стрит, Фрэнк обнаруживает, что мешок шевелится. Впервые за все время Мартин нарушает правило, заглядывает внутрь пакета и обнаруживает там красивую женщину, которая оказывается дочерью видного китайского мафиози...', '2002-10-02', 'film'),
  ('/static/movies/transformers/movie.mp4', '/static/movies/transformers/logo.png', 'Трансформеры', 'Сэм хочет крутую тачку, а получает живого робота. Начало масштабной франшизы Майкла Бэя о машинах и людях', '/static/movies/transformers/card.png', '/static/movies/transformers/poster.png', 2, 7.6, 'В течение многих столетий две расы роботов-инопланетян — Автоботы и Десептиконы — вели войну, ставкой в которой была судьба Вселенной. И вот война докатилась до Земли. В то время, когда силы зла ищут ключ к верховной власти, наш последний шанс на спасение находится в руках юного землянина. Единственное, что стоит между несущими зло Десептиконами и высшей властью - это ключ, находящийся в руках простого парнишки. Обычный подросток, Сэм Уитвикки озабочен повседневными хлопотами — школа, друзья, машины, девочки. Не ведая о том, что он является последним шансом человечества на спасение, Сэм вместе со своей подругой Микаэлой оказывается вовлеченным в противостояние Автоботов и Десептиконов. Только тогда Сэм понимает истинное значение семейного девиза Уитвикки — «без жертв нет победы!»', '2007-07-04', 'film'),
  ('/static/movies/wolf-of-wall-street/movie.mp4', '/static/movies/wolf-of-wall-street/logo.png', 'Волк с Уолл-Стрит', 'Восхождение циника-гедониста на бизнес-олимп 1980-х. Блистательный тандем Леонардо ДиКаприо и Мартина Скорсезе', '/static/movies/wolf-of-wall-street/card.png', '/static/movies/wolf-of-wall-street/poster.png', 2, 8.0, '1987 год. Джордан Белфорт становится брокером в успешном инвестиционном банке. Вскоре банк закрывается после внезапного обвала индекса Доу-Джонса. По совету жены Терезы Джордан устраивается в небольшое заведение, занимающееся мелкими акциями. Его настойчивый стиль общения с клиентами и врождённая харизма быстро даёт свои плоды. Он знакомится с соседом по дому Донни, торговцем, который сразу находит общий язык с Джорданом и решает открыть с ним собственную фирму. В качестве сотрудников они нанимают нескольких друзей Белфорта, его отца Макса и называют компанию «Стрэттон Оукмонт». В свободное от работы время Джордан прожигает жизнь: лавирует от одной вечеринки к другой, вступает в сексуальные отношения с проститутками, употребляет множество наркотических препаратов, в том числе кокаин и кваалюд. Однажды наступает момент, когда быстрым обогащением Белфорта начинает интересоваться агент ФБР...', '2013-12-09', 'film'),

  ('/static/movies/avengers/movie.mp4', '/static/movies/avengers/logo.png', 'Мстители', 'Титан Танос вынашивает страшный план — угрозу всей Вселенной. Предпоследний фильм о суперкоманде Marvel', '/static/movies/avengers/card.png', '/static/movies/avengers/poster.png', 2, 7.9, 'Пока Мстители и их союзники продолжают защищать мир от различных опасностей, с которыми не смог бы справиться один супергерой, новая угроза возникает из космоса: Танос. Межгалактический тиран преследует цель собрать все шесть Камней Бесконечности - артефакты невероятной силы, с помощью которых можно менять реальность по своему желанию. Всё, с чем Мстители сталкивались ранее, вело к этому моменту – судьба Земли никогда ещё не была столь неопределённой.', '2012-04-11', 'film'),
  ('/static/movies/drunk/movie.mp4', '/static/movies/drunk/logo.png', 'Еще по одной', 'Четыре школьных учителя проверяют на себе гипотезу о пользе алкоголя. Душеспасительное драмеди с Мадсом Миккельсеном', '/static/movies/drunk/card.png', '/static/movies/drunk/poster.png', 7, 7.6, 'Фильм получил премию «Оскар» в номинации «Лучший фильм на иностранном языке». Трогательная драма о том, что будет, если выпивать всегда, но по чуть-чуть. Школьный учитель Мартин (Мадс Миккельсен) испытывает выгорание: ученики не стараются, дни тянутся, словно вагоны нескончаемого поезда. Вместе с друзьями он решает проверить одну теорию. Всем известно, как хорошо бывает после первого бокала. А что, если постоянно и равномерно поддерживать градус, не напиваясь? Задумка быстро приносит положительные результаты, мир расцветает. Но долго ли это продлится? Режиссер фильмов «Курск» и «Охота» представляет историю, в которой нетрудно будет узнать самих себя.', '2020-09-12', 'film'),
  ('/static/movies/ford-v-ferrari/movie.mp4', '/static/movies/ford-v-ferrari/logo.png', 'Форд против Феррари', 'Автоконструктор и строптивый гонщик против непобедимых конкурентов. Биографическая драма о любви к скорости', '/static/movies/ford-v-ferrari/card.png', '/static/movies/ford-v-ferrari/poster.png', 2, 8.2, 'В начале 1960-х Генри Форд II принимает решение улучшить имидж компании и сменить курс на производство более модных автомобилей. После неудавшейся попытки купить практически банкрота Ferrari американцы решают бросить вызов итальянским конкурентам на трассе и выиграть престижную гонку 24 часа Ле-Мана. Чтобы создать подходящую машину, компания нанимает автоконструктора Кэррола Шэлби, а тот отказывается работать без выдающегося, но, как считается, трудного в общении гонщика Кена Майлза. Вместе они принимаются за разработку впоследствии знаменитого спорткара Ford GT40.', '2019-08-30', 'film'),
  ('/static/movies/greenbook/movie.mp4', '/static/movies/greenbook/logo.png', 'Зеленая книга', 'Простой белый парень сопровождает известного чернокожего музыканта во время тура по югу США в шестидесятые годы', '/static/movies/greenbook/card.png', '/static/movies/greenbook/poster.png', 2, 8.5, 'Оскароносное роуд-муви о дружбе утонченного музыканта и простого парня из Бронкса. Реальная история джазового пианиста Дона Ширли и его хамоватого водителя Тони. Несмотря на свою непохожесть, герои постепенно становятся верными друзьями. Кино о дружбе, которая ломает стереотипы и расовые предрассудки. О человечности и ошибках прошлых поколений, которые не стоит повторять.', '2018-09-11', 'film'),
  ('/static/movies/once-in-hollywood/movie.mp4', '/static/movies/once-in-hollywood/logo.png', 'Однажды в Голливуде', 'Можно ли переписать историю? Самый ностальгический фильм Тарантино — с Шэрон Тейт, Брюсом Ли и Чарли Мэнсоном', '/static/movies/once-in-hollywood/card.png', '/static/movies/once-in-hollywood/poster.png', 2, 5.9, 'Фильм повествует о череде событий, произошедших в Голливуде в 1969 году, на закате его «золотого века». Известный актер Рик Далтон и его дублер Клифф Бут пытаются найти свое место в стремительно меняющемся мире киноиндустрии.', '2008-01-19', 'film'),
  ('/static/movies/lamborghini/movie.mp4', '/static/movies/lamborghini/logo.png', 'Ламборгини', 'Ферруччо Ламборгини собирает красивый спорткар, который утрет нос концерну Феррари. Байопик с Фрэнком Грилло', '/static/movies/lamborghini/card.png', '/static/movies/lamborghini/poster.png', 2, 6.7, 'Ферруччо Ламборгини (Романо Реджиани) возвращается домой с полей Второй мировой. Италия находится в руинах. Но если для кого-то это трагедия, то для других возможности. Ламборгини не собирается копаться в земле, как его мама (Франческа Де Мартини) и папа (Фортунато Серлино). Он хочет собирать спорткары. Всю жизнь он интересовался различными механизмами и металлическими конструкциями. Молодой человек верит в себя, и ему удается построить успешный бизнес по конструированию тракторов, но свою мечту Ферруччо не оставляет. Уже во взрослом возрасте Ламборгини (Фрэнк Грилло) собирает неземной красоты спорткар, надеясь обогнать конкурента в лице Энцо Феррари (Гэбриел Бирн). Байопик знаменитого итальянского инженера и бизнесмена снял Бобби Мореско, обладатель «Оскара» за лучший оригинальный сценарий «Столкновения» и один из продюсеров «Малышки на миллион» Клинта Иствуда.', '2022-10-23', 'film'),
  ('/static/movies/legend/movie.mp4', '/static/movies/legend/logo.png', 'Легенда', 'Взлёт и падение легендарных близнецов Крэй, главных преступников 60-х годов', '/static/movies/legend/card.png', '/static/movies/legend/poster.png', 3, 7.2, 'Байопик о легендарный лондонских генгстерах Крэй и их преступной карьере в 1960-е. Страдающий психическим расстройством и склонный к насилию Рональд по решению суда помещён в специализированную клинику, но его брат Реджи надавил на врача и преступник оказался на свободе, чтобы помочь близнецу стать главными мафиози Британии. Официально Крэй владеют ночным клубом, а в свободное время практикуют рэкет, убийства и грабежи. Реджинальд вступает в романтические отношения с юной Фрэнсис и та рассчитывает, что ради любви, гангстер завяжет с криминалом. Но захочет ли очаровательный глава преступности Ист-Энда встать на путь исправления?"', '2015-09-03', 'film'),
  ('/static/movies/pele/movie.mp4', '/static/movies/pele/logo.png', 'Пеле: Рождение легенды', 'Героическая история о начале карьеры лучшего футболиста XX века', '/static/movies/pele/card.png', '/static/movies/pele/poster.png', 2, 8.4, 'Футболист Пеле стал суперзвездой, выступив на чемпионате мира 1958 года. После он профессионально играл в Бразилии ещё два десятилетия, выиграв ещё три главных кубков. В 1999 году он был назван «игроком века» по версии ФИФА. Как начинал свой путь человек, историю которого можно смело назвать не биографией, а агиографией — жизнь бога футбола? Кто научил босоногого мальчишку невероятным приёмам и поддержал в его стремлении прославить Родину?', '2016-04-23', 'film'),
  ('/static/movies/streltsov/movie.mp4', '/static/movies/streltsov/logo.png', 'Стрельцов', 'Испытания всесоюзной славой, громкие скандалы и тюрьма. Александр Петров в роли самого свободолюбивого футболиста', '/static/movies/streltsov/card.png', '/static/movies/streltsov/poster.png', 1, 7.1, 'Отечественная спортивная мелодрама об одном из величайших футболистов Советского Союза. К 20 годам Эдуарда Стрельцова знает весь Союз. Он стремительно ворвался в сборную, хотя ещё пару лет назад играл за любительскую команду завода. Сегодня же он лучший игрок «Торпедо», ему доверяет опытный тренер Виктор Маслов. Стрельцов ходит в дорогие рестораны, любит потанцевать и посидеть с друзьями. Режим удаётся соблюдать далеко не всегда, но улыбчивому парню прощают очень многое — такой талант нельзя загубить. В ближайшее время сборная СССР должна отправиться в Швецию на чемпионат мира. В стране ожидают схватки Стрельцова с другим знаменитым юным футболистом — Пеле. Прямо перед отъездом на ответственный турнир торпедовца обвиняют в изнасиловании. Суд приговаривает его к 12 годам заключения. Он отсидит 5 лет, после чего найдёт в себе силы вернуться в большой футбол и вновь напомнить о себе.', '2020-09-24', 'film'),
  ('/static/movies/wrath-of-man/movie.mp4', '/static/movies/wrath-of-man/logo.png', 'Гнев человеческий', 'Нелюдимый парень с таинственным прошлым устраивается на работу инкассатором, чтобы совершить священный акт мести', '/static/movies/wrath-of-man/card.png', '/static/movies/wrath-of-man/poster.png', 3, 7.6, 'В инкассаторскую компанию устраивается новый сотрудник Эйч (Джейсон Стэйтем). Он тих и немногословен, что, в принципе, приветствуется в такой сфере, где каждый день перевозят по Лос-Анджелесу миллионы долларов наличными. Но внутри этого загадочного британца бушуют страсти. Он жаждет отомстить банде грабителей, нападавших на инкассаторские машины. Это история мести, которая перемещается по временной шкале то в прошлое, то в настоящее, а также показанная с позиции разных персонажей.', '2021-04-22', 'film'),
  ('/static/movies/brat2/movie.mp4', '/static/movies/brat2/logo.png', 'Брат 2', 'Американцы знакомятся с Данилой Багровым и узнают, в чем сила. Сиквел о герое времени с мощным рок-саундтреком', '/static/movies/brat2/card.jpg', '/static/movies/brat2/poster.png', 1, 8.2, 'Сиквел культового боевика о похождениях Данилы Багрова. Прошло два года после событий предыдущего фильма, Данила Багров ищет своё счастье в Москве. Получив приглашение поучаствовать в съемках программы, посвященной героям Чеченской войны, он встречается со своими старыми товарищами. Но вскоре одного из них настигает бандитская пуля, и Данила оказывается втянутым в новые разборки, которые на этот раз заведут его в далекую Америку.', '2000-08-30', 'film');

INSERT INTO public.movie_genres (movie_id, genre_id)
VALUES
  (1, 1),
  (1, 3),
  (1, 5),
  (2, 1),
  (2, 4),
  (3, 4),
  (3, 5),
  (4, 3),
  (5, 6),
  (5, 4),
  (5, 3),
  (6, 2),
  (7, 2),
  (7, 5),
  (8, 4),
  (8, 8),
  (9, 1),
  (9, 3),
  (9, 5),
  (10, 1),
  (10, 4),
  (11, 4),
  (11, 5),
  (12, 3),
  (13, 6),
  (13, 4),
  (13, 3),
  (14, 2),
  (15, 2),
  (15, 5),
  (16, 4),
  (16, 8),
  (17, 1),
  (17, 4),
  (18, 4),
  (18, 5),
  (19, 3),
  (20, 6),
  (20, 4),
  (20, 3),
  (21, 2),
  (22, 2),
  (22, 5),
  (23, 4),
  (24, 8),
  (25, 8),
  (26, 8),
  (27, 8);

INSERT INTO public.collections (title, card_url)
VALUES
  ('Лучшие сериалы', ''),
  ('Классика 10-х годов', ''),
  ('Выбор редакции', '');

INSERT INTO public.movie_collections (movie_id, collection_id)
VALUES
  (1, 1),
  (2, 1),
  (3, 1),
  (4, 2),
  (5, 2),
  (6, 2),
  (7, 2),
  (8, 2),
  (9, 2),
  (10, 2),
  (11, 2),
  (12, 2),
  (13, 2),
  (14, 2),
  (15, 2),
  (16, 2),
  (17, 3),
  (18, 3),
  (19, 3),
  (20, 3),
  (21, 3),
  (22, 3),
  (23, 3),
  (24, 3),
  (25, 3),
  (26, 3),
  (27, 3);

INSERT INTO public.movie_staff (first_name, second_name, patronymic, biography, country_id, post)
VALUES
  ('Райан', 'Гослинг', '', 'Снимался в барби', 2, 'actor'),
  ('Кристофер', 'Нолан', '', 'Крутой', 2, 'director'),
  ('Александр', 'Петров', 'Федорович', 'Лучший русский актер', 1, 'actor');

INSERT INTO public.staff_members (movie_id, movie_staff_id)
VALUES
  (1, 1),
  (1, 3),
  (1, 2),
  (2, 1),
  (2, 2),
  (3, 2),
  (3, 3),
  (4, 3),
  (5, 3),
  (6, 2),
  (7, 3);
