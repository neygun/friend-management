CREATE TABLE IF NOT EXISTS user (
    id BIGINT PRIMARY KEY,
    email VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT USER_EMAIL_UQ UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS relationship (
    id BIGINT PRIMARY KEY,
    requestor_id BIGINT NOT NULL,
    target_id BIGINT NOT NULL,
    "type" VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT RELATIONSHIP_REQUESTOR_ID_FK FOREIGN KEY (requestor_id) REFERENCES user(id),
    CONSTRAINT RELATIONSHIP_TARGET_ID_FK FOREIGN KEY (target_id) REFERENCES user(id),
    CONSTRAINT RELATIONSHIP_REQUESTOR_ID_TARGET_ID_TYPE_UQ UNIQUE (requestor_id, target_id, "type")
);
