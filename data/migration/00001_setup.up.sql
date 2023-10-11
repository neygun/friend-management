CREATE TABLE IF NOT EXISTS USER (
    ID BIGINT PRIMARY KEY,
    EMAIL VARCHAR NOT NULL,
    CREATED_AT TIMESTAMPTZ NOT NULL,
    UPDATED_AT TIMESTAMPTZ  NOT NULL,
    CONSTRAINT USER_EMAIL_UQ UNIQUE (EMAIL)
);

CREATE TABLE IF NOT EXISTS RELATIONSHIP (
    ID BIGINT PRIMARY KEY,
    REQUESTOR_ID BIGINT NOT NULL,
    TARGET_ID BIGINT NOT NULL,
    "TYPE" VARCHAR NOT NULL,
    CREATED_AT TIMESTAMPTZ NOT NULL,
    UPDATED_AT TIMESTAMPTZ  NOT NULL,
    CONSTRAINT RELATIONSHIP_REQUESTOR_ID_FK FOREIGN KEY (REQUESTOR_ID) REFERENCES USER(ID),
    CONSTRAINT RELATIONSHIP_TARGET_ID_FK FOREIGN KEY (TARGET_ID) REFERENCES USER(ID),
    CONSTRAINT RELATIONSHIP_REQUESTOR_ID_TARGET_ID_TYPE_UQ UNIQUE (REQUESTOR_ID, TARGET_ID, "TYPE")
);
