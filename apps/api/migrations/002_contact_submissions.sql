-- Submissions from the home-page contact form. source_ip and user_agent
-- captured for moderation; handled_at flips when Cole replies (manual at v1).
CREATE TABLE IF NOT EXISTS contact_submissions (
    id           uuid        PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at   timestamptz NOT NULL DEFAULT now(),
    name         text        NOT NULL,
    email        text        NOT NULL,
    message      text        NOT NULL,
    source_ip    inet,
    user_agent   text,
    handled_at   timestamptz
);

CREATE INDEX IF NOT EXISTS idx_contact_submissions_created_at
    ON contact_submissions (created_at DESC);
