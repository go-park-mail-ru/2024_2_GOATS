wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

-- Функция для генерации уникального хэша
function generate_unique_hash(thread_id)
  local characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  local length = math.random(5, 15)
  local hash = thread_id .. os.time()

  for i = 1, length do
    local random_index = math.random(1, #characters)
    local random_char = characters:sub(random_index, random_index)
    hash = hash .. random_char
  end

  return hash
end

-- Инициализация потоков
threads = {}

function setup(thread)
  thread:set("thread_id", #threads + 1) -- Уникальный ID потока
  table.insert(threads, thread)
end

max_requests = 100000
-- Генерация запроса
request = function()
  local unique_hash = generate_unique_hash(thread_id)
  local username = "user_" .. unique_hash
  local email = "user" .. unique_hash .. "@example.com"
  local password_hash = "hashed_password_" .. unique_hash
  local birthdate = "1990-01-" .. math.random(1, 28)

  local body = string.format(
    '{"username":"%s","email":"%s","password":"%s","passwordConfirmation":"%s"}',
    username,
    email,
    password_hash,
    password_hash
  )

  return wrk.format(nil, nil, nil, body)
end
